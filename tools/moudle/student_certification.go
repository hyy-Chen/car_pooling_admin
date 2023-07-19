/**
 @author : ikkk
 @date   : 2023/7/19
 @text   : nil
**/

package moudle

import "time"

type StudentCertification struct {
	Id           int64 `gorm:"primaryKey"`
	UserId       int64
	UserOpenid   string
	UserSchoolId string
	AuditorId    int64
	UserPic      string
	AdminInputId string
	State        int
	Message      string
	Deleted      int
	CreateTime   time.Time
	UpdateTime   time.Time
	// 关联到user表，之后进行关联查询以及修改
	User User
}

func (StudentCertification) TableName() string {
	return "student_certification"
}
