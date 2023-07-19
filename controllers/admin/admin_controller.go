/**
 @author : ikkk
 @date   : 2023/7/16
 @text   : nil
**/

package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminController struct {
}

func (adm AdminController) Ok(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
