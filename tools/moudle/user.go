/**
 @author : ikkk
 @date   : 2023/7/19
 @text   : nil
**/

package moudle

import "time"

type User struct {
	Id                int64 `gorm:"primaryKey"`
	SchoolId          string
	Name              string
	NickName          string
	Phone             string
	Sex               int8
	Account           string
	SuccessesNumber   int
	TotalNumber       int
	ReliabilityRating int
	WaitTime          int
	Openid            string
	SessionKey        string
	Deleted           int
	CreateTime        time.Time
	UpdateTime        time.Time
}

func (User) TableName() string {
	return "user"
}
