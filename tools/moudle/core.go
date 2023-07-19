/**
 @author : ikkk
 @date   : 2023/7/19
 @text   : 相关gorm操作查看 https://gorm.io/zh_CN/docs/create.html
**/

package moudle

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:F71BBE7314684A96@tcp(8.134.152.127:3306)/pooling_car?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
