/**
 @author : ikkk
 @date   : 2023/7/24
 @text   : nil
**/

package admin_controllers

import (
	"car_pooling_admini/tools/moudle"
	"car_pooling_admini/tools/tools"
	Result "car_pooling_admini/tools/tools/gin_result"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type Admin struct {
}

// GetAdminInformation 获取用户信息接口
func (Admin) GetAdminInfo(c *gin.Context) {
	// 获取id
	adminId, _, _ := tools.ParseToken(c.GetHeader("Authorization"))
	// 得到用户信息
	sysUser, err := moudle.GetSysUserById(adminId)
	// 判别错误类型
	if err == gorm.ErrRecordNotFound {
		c.JSON(400, Result.ErrAdminAccount)
	} else if err == nil {
		sysUser.Password = ""
		c.JSON(200, Result.OK.WithData(sysUser))
	} else {
		c.JSON(500, Result.Err)
	}
}

type PutAdminInformationRequest struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Password string `json:"password"`
}

// PutAdminInfo 更新账号信息
func (Admin) PutAdminInfo(c *gin.Context) {
	// 获取id
	adminId, _, _ := tools.ParseToken(c.GetHeader("Authorization"))
	// 得到用户信息
	sysUser, err := moudle.GetSysUserById(adminId)
	if err == gorm.ErrRecordNotFound {
		c.JSON(400, Result.ErrAdminAccount)
	} else if err == nil {
		// 更新信息
		// 解析前端传回信息
		// 解析JSON数据
		user := PutAdminInformationRequest{}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, Result.ErrParamType)
			return
		}
		user.Id = 0
		if user.Password != "" {
			user.Password = tools.PasswordEncryption(user.Password)
		}
		// 存储回数据库
		// 删除缓存key
		key := moudle.SysUserInformationKeyPrefix + strconv.FormatInt(sysUser.Id, 10)
		moudle.Cli.Del(moudle.Ctx, key)
		// 更新数据库
		moudle.DB.Model(&sysUser).Updates(user).Scan(&sysUser)
		// 更新redis
		moudle.Cli.Del(moudle.Ctx, key)
		moudle.Cli.Set(moudle.Ctx, key, sysUser, moudle.SysUserInformationKeyTime)
		sysUser.Password = ""
		c.JSON(200, Result.OK.WithData(sysUser))
	} else {
		c.JSON(500, Result.Err)
	}
}
