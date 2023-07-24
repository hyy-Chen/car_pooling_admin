/**
 @author : ikkk
 @date   : 2023/7/19
 @text   : nil
**/

package moudle

import (
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"strconv"
	"time"
)

type SysUser struct {
	Id           int64                 `json:"id" gorm:"primaryKey"`
	SchoolId     string                `json:"schoolId"`
	Name         string                `json:"name"`
	NickName     string                `json:"nickName"`
	Account      string                `json:"account"`
	Password     string                `json:"password,omitempty"`
	CreatePerson int64                 `json:"createPerson"`
	Role         int                   `json:"role"`
	Deleted      soft_delete.DeletedAt `json:"-" gorm:"softDelete:flag"`
	CreateTime   time.Time             `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime   time.Time             `json:"updateTime" gorm:"autoUpdateTime"`
}

var SysUserInformationKeyTime = time.Hour * 14 * 24
var SysUserInformationKeyPrefix = "admin_controllers:info:id:"
var SysUserTokenKeyTime = time.Hour * 24 * 14
var SysUserTokenKeyPrefix = "admin_controllers:token:id:"
var SysUserTokenTTLKeyPrefix = "admin_controllers:token:expiration:"

func (SysUser) TableName() string {
	return "sys_user"
}

func (s *SysUser) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SysUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

// GetSysUserById 通过id查询用户信息, 先从缓存查找，缓存不存在再去数据库查找
func GetSysUserById(id string) (*SysUser, error) {
	// 先查询缓存
	result, err := Cli.Get(Ctx, SysUserInformationKeyPrefix+id).Result()
	if err == redis.Nil {
		// 去数据库查找
		user := SysUser{}
		err1 := DB.Where("id=?", id).First(&user).Error
		if err1 == gorm.ErrRecordNotFound {
			return nil, err1
		} else {
			// 将得到的信息存储回redis里, 这里如果出现错误也暂时不解决, 打印到日志先
			Cli.Set(Ctx, SysUserInformationKeyPrefix+id, &user, SysUserInformationKeyTime)
			return &user, nil
		}
	} else {
		// 反序列化再返回
		user := SysUser{}
		err2 := json.Unmarshal([]byte(result), &user)
		return &user, err2
	}
}

// GetSysUserByAccount 通过id查询用户信息 去数据库查找, 查找到后更新缓存信息
func GetSysUserByAccount(account string) (*SysUser, error) {
	// 去数据库查找
	user := SysUser{}
	err := DB.Where("account=?", account).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	} else {
		Cli.Set(Ctx, SysUserInformationKeyPrefix+strconv.FormatInt(user.Id, 10), &user, SysUserInformationKeyTime)
		return &user, nil
	}
}
