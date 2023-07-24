/**
 @author : ikkk
 @date   : 2023/7/19
 @text   : nil
**/

package moudle

import (
	"encoding/json"
	"time"
)

type SysUser struct {
	Id           int64 `gorm:"primaryKey"`
	SchoolId     string
	Name         string
	NickName     string
	Account      string
	Password     string
	CreatePerson int64
	Role         int
	Deleted      int
	CreateTime   time.Time
	UpdateTime   time.Time
}

func (SysUser) TableName() string {
	return "sys_user"
}

func (s *SysUser) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SysUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
