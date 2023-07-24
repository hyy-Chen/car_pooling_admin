/**
 @author : ikkk
 @date   : 2023/7/16
 @text   : nil
**/

package api

import (
	"car_pooling_admini/tools/moudle"
	"car_pooling_admini/tools/tools"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type AdminController struct {
}

func (AdminController) Ok(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// 登录api
type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// Login 登录操作
// 先获取账号密码，之后去数据库内查找对应的账号密码是否存在，查找后获取账号信息，将id通过jwt打包成token再发送给前端
// 缓存存储id->账号信息
func (*AdminController) LoginHandler(c *gin.Context) {

	// 解析请求中的JSON数据到LoginRequest结构体
	loginRequest := LoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// 数据判空
	if loginRequest.Account == "" || loginRequest.Password == "" {
		c.JSON(401, gin.H{"status": "error", "message": "账号或者密码不能为空"})
		return
	}
	// 去数据库进行搜索，得到结果
	sysUser := moudle.SysUser{}

	// 进行账号进去数据库查找
	tx := moudle.DB.Where("account=? and deleted=?", loginRequest.Account, 0).First(&sysUser)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		c.JSON(401, gin.H{"status": "error", "message": "账号不存在"})
		return
	}
	err := tools.PasswordValidation(loginRequest.Password, sysUser.Password)
	if err != nil {
		c.JSON(401, gin.H{"status": "error", "message": "密码错误"})
		return
	}
	// 找到了符合条件的数据，将id->结构体进行redis存储
	set := moudle.Cli.Set(moudle.Ctx, strconv.FormatInt(sysUser.Id, 10), &sysUser, 0) // 现在展示设置为0，之后调整时间
	fmt.Println(set.Result())
	// 获得token, 默认设置2天
	token, err := tools.GenerateToken(strconv.FormatInt(sysUser.Id, 10))
	if err != nil {
		c.JSON(401, gin.H{"status": "error", "message": "异常错误，请联系管理员查看后台"})
		return
	}
	// 保存token在redis里, 时间定为14天
	tokenInRedisTime := 14 * 24 * time.Hour
	moudle.Cli.Set(moudle.Ctx, token, "", tokenInRedisTime)

	// 返回登录成功, 将token返回
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
	return
}
