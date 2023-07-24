/**
 @author : ikkk
 @date   : 2023/8/9
 @text   : nil
**/

package routers

import (
	"car_pooling_admini/controllers/admin_controllers"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	adminRouters.Use(authorizationInterceptor())
	{
		adminRouters.GET("/", admin_controllers.Admin{}.GetAdminInfo)
		adminRouters.PUT("/", admin_controllers.Admin{}.PutAdminInfo)
	}
	adminInfoRouters := adminRouters.Group("/info")
	adminInfoRouters.Use(detectAdminStatusInterceptor())
	{
		adminInfoRouters.GET("/")
	}
}
