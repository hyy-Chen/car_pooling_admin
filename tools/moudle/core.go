/**
 @author : ikkk
 @date   : 2023/7/19
 @text   : 相关gorm操作查看 https://gorm.io/zh_CN/docs/create.html
**/

package moudle

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
	Ctx = context.Background()
	Cli *redis.Client
)

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:F71BBE7314684A96@tcp(8.134.152.127:3306)/pooling_car?charset=utf8mb4&parseTime=True&loc=Local"
	// 获取连接
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 获取redis连接
	Cli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址和端口
		Password: "",               // 服务器密码（如果没有设置密码，保持为空）
		DB:       0,                // 使用默认数据库
	})
	// 测试连接是否成功
	_, err = Cli.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}
}
