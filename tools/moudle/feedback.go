/**
 @author : ikkk
 @date   : 2023/7/19
 @text   : nil
**/

package moudle

import "time"

type Feedback struct {
	Id         int64 `gorm:"primaryKey"`
	AuditorId  int64
	Question   string
	Reply      string
	Accepted   int
	Deleted    int
	CreateTime time.Time
	UpdateTime time.Time
}

func (Feedback) TableName() string {
	return "feedback"
}
