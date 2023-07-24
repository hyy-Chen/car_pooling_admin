/**
 @author : ikkk
 @date   : 2023/8/10
 @text   : nil
**/

package routers

import (
	"car_pooling_admini/tools/moudle"
	"car_pooling_admini/tools/tools"
	Result "car_pooling_admini/tools/tools/gin_result"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"time"
)

// 令牌拦截器中间件, 用于验证操作者是否有令牌以及令牌是否过期等等操作
func authorizationInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 进行令牌验证，令牌校验不过关的话进行拦截
		// 获取令牌
		token := c.GetHeader("Authorization")
		// 判断令牌是否为空
		if token == "" {
			c.JSON(400, Result.ErrTokenLack)
			c.Abort()
			return
		}
		// 验证令牌是否还存在于redis里，如果不存在就直接返回错误
		exists := moudle.Cli.Exists(moudle.Ctx, moudle.SysUserTokenTTLKeyPrefix+token).Val()
		if exists == 0 {
			c.JSON(400, Result.ErrTokenInvalidated)
			c.Abort()
			return
		}
		// 如果令牌存在就检查token里过期时间是否过期，如果token里过期就更新返回
		// 验证令牌是否可用，是否需要更新
		id, t, err := tools.ParseToken(token)
		if err != nil {
			c.JSON(400, Result.ErrTokenIllegal)
			c.Abort()
			return
		}
		// 判断令牌是否过期
		if t.Unix() < time.Now().Unix() {
			// 如果令牌已经过期
			//得到新的令牌
			generateToken, err := tools.GenerateToken(id)
			if err != nil {
				log.Printf("生成Token失败: %v", err)
				c.JSON(500, Result.Err)
				c.Abort()
				return
			}
			// 令牌未过期更新令牌以及缓存，并且要求前端更新令牌
			// 1.在redis里去除旧的令牌
			// 2.在redis里去除id->token映射
			// 3.重新设置令牌
			// 4.重新设置id->token映射
			_, err = moudle.Cli.TxPipelined(moudle.Ctx, func(pipe redis.Pipeliner) error {
				//在redis里去除旧的令牌
				pipe.Del(moudle.Ctx, moudle.SysUserTokenTTLKeyPrefix+token)
				// 将token存储
				pipe.Set(moudle.Ctx, moudle.SysUserTokenTTLKeyPrefix+generateToken, "", moudle.SysUserTokenKeyTime)
				// 存储用户id->token
				pipe.Set(moudle.Ctx, moudle.SysUserTokenKeyPrefix+id, generateToken, moudle.SysUserTokenKeyTime)
				return nil
			})
			if err != nil {
				// 服务器错误
				log.Printf("redis操作失败: %v", err)
				c.JSON(500, Result.Err)
				c.Abort()
				return
			}
			// 将新的token返回，要求前端更新
			c.JSON(401, Result.ErrTokenExpires.WithData(gin.H{"token": token}))
			c.Abort()
			return
		}
	}
}

// 用户信息拦截器中间件，
// 确保管理员权限存在的拦截器（预防管理员权限不足但是调用了对应方法的情况）
func detectAdminStatusInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取令牌信息，经过第一个检测得到的令牌可以用
		token := c.GetHeader("Authorization")
		// 获取到用户id
		id, _, _ := tools.ParseToken(token)
		user, err := moudle.GetSysUserById(id)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 如果此时有err，必然是用户被删除后使用令牌登录,那么销毁令牌防止重复登录消耗资源
				moudle.Cli.Del(moudle.Ctx, moudle.SysUserTokenTTLKeyPrefix+token)
				c.JSON(400, Result.ErrTokenAccount)
			} else {
				c.JSON(500, Result.Err)
			}
			c.Abort()
			return
		}
		// 用户权限不足警告
		if user.Role != 1 {
			c.JSON(400, Result.ErrTokenRole)
			c.Abort()
			return
		}
		c.Next()
		return
	}
}
