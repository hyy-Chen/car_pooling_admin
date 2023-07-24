/**
 @author : ikkk
 @date   : 2023/7/16
 @text   : nil
**/

package routers

import (
	"car_pooling_admini/controllers/default_controllers"
	"github.com/gin-gonic/gin"
)

// DefaultRoutersInit 初始化路由端口并配置对应方法
func DefaultRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/")
	{
		apiRouters.GET("/ok", default_controllers.OK)
		apiRouters.POST("/login", default_controllers.LoginHandler)
	}
}
