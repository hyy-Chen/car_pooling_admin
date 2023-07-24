/**
 @author : ikkk
 @date   : 2023/8/9
 @text   : nil
**/

package default_controllers

import (
	Result "car_pooling_admini/tools/tools/gin_result"
	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context) {
	c.JSON(200, Result.OK)
}
