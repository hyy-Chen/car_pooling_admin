/**
 @author : ikkk
 @date   : 2023/7/16
 @text   : nil
**/

package routers

import (
	"car_pooling_admini/controllers/admin"
	"github.com/gin-gonic/gin"
)

// AdminRoutersInit 初始化路由端口并配置对应方法
func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/ok", admin.AdminController{}.Ok)
	}
}
