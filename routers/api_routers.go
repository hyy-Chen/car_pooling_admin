/**
 @author : ikkk
 @date   : 2023/7/16
 @text   : nil
**/

package routers

import (
	"car_pooling_admini/controllers/api"
	"car_pooling_admini/tools/moudle"
	"car_pooling_admini/tools/tools"
	"github.com/gin-gonic/gin"
	"time"
)

// 拦截器中间件, 用于验证操作者是否有令牌已经令牌是否过期等等操作
func authorizationInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前的 API 路径
		apiPath := c.Request.URL.Path
		// 对于登录以及测试api, 不进行令牌验证
		switch apiPath {
		case "/login":
		case "ok":
			c.Next()
			return
		}
		// 进行令牌验证，令牌校验不过关的话进行拦截
		// 获取令牌
		token := c.GetHeader("Authorization")
		// 判断令牌是否为空
		if token == "" {
			c.JSON(401, gin.H{"status": "error", "message": "The token is empty"})
			return
		}
		// 验证令牌是否可用，是否需要更新
		id, t, err := tools.ParseToken(token)
		// 如果令牌被篡改
		if err != nil {
			c.JSON(401, gin.H{"status": "error", "message": "Token error"})
			return
		}
		// 判断令牌是否过期
		if t.Unix() < time.Now().Unix() {
			// 判断令牌是否在redis里过期
			result, err := moudle.Cli.Exists(moudle.Ctx, token).Result()
			// 如果令牌在redis里也已经过期，就要要求重新登录
			if err != nil || result == 0 {
				c.JSON(401, gin.H{"status": "error", "message": "Token expiration"})
				return
			}
			// 令牌未过期更新令牌以及缓存，并且要求前端更新令牌
			// 在redis里去除旧的令牌
			moudle.Cli.Del(moudle.Ctx, token)
			// 得到新的令牌
			generateToken, err := tools.GenerateToken(id)
			if err != nil {
				// 之后看看如何操作
			}
			// 重新设置redis里token
			moudle.Cli.Set(moudle.Ctx, generateToken, "", 14*24*time.Hour)
			// 将新的token返回，要求前端更新
			c.JSON(402, gin.H{"status": "fail", "message": "Token updates", "data": gin.H{"token": token}})
			return
		}
	}
}

// AdminRoutersInit 初始化路由端口并配置对应方法
func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/api")
	// 应用拦截器
	adminRouters.Use(authorizationInterceptor())
	{
		adminRouters.GET("/ok", api.AdminController{}.Ok)
		adminRouters.POST("/login", api.AdminController{}.LoginHandler)
	}
}
