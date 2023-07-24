/**
 @author : ikkk
 @date   : 2023/7/16
 @text   : nil
**/

package default_controllers

import (
	"car_pooling_admini/tools/moudle"
	"car_pooling_admini/tools/tools"
	Result "car_pooling_admini/tools/tools/gin_result"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

// 登录api
type loginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// LoginHandler 登录操作
// 先获取账号密码，之后去数据库内查找对应的账号密码是否存在，查找后获取账号信息，将id通过jwt打包成token再发送给前端
// 缓存存储id->账号信息
func LoginHandler(c *gin.Context) {
	// 解析请求中的JSON数据到LoginRequest结构体
	loginRequest := loginRequest{}
	// 解析错误，内容格式不对
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, Result.ErrParamType)
		return
	}
	// 数据判空
	if loginRequest.Account == "" || loginRequest.Password == "" {
		c.JSON(400, Result.ErrParamMiss)
		return
	}
	// 去数据库进行搜索，得到结果
	sysUser, err := moudle.GetSysUserByAccount(loginRequest.Account)
	// 数据库不存在信息
	if err != nil {
		c.JSON(400, Result.ErrLoginAccount)
	}
	// 验证密码
	err = tools.PasswordValidation(loginRequest.Password, sysUser.Password)
	if err != nil {
		c.JSON(400, Result.ErrLoginPassword)
		return
	}
	id := strconv.FormatInt(sysUser.Id, 10)
	// 获得token, 默认设置7天
	token, err := tools.GenerateToken(id)
	if err != nil {
		log.Printf("生成Token失败: %v", err)
		c.JSON(500, Result.Err)
		return
	}
	// 获取旧的token
	oldToken, err := moudle.Cli.Get(moudle.Ctx, moudle.SysUserTokenKeyPrefix+id).Result()
	// 找到了符合条件的数据，将id->结构体进行redis存储 key : "adminId:" + id
	pipe := moudle.Cli.TxPipeline()
	// 删除旧token
	if err != redis.Nil {
		pipe.Del(moudle.Ctx, moudle.SysUserTokenTTLKeyPrefix+oldToken)
	}
	// 存储用户信息
	pipe.Set(moudle.Ctx, moudle.SysUserInformationKeyPrefix+id, sysUser, moudle.SysUserInformationKeyTime)
	// 将token存储
	pipe.Set(moudle.Ctx, moudle.SysUserTokenTTLKeyPrefix+token, "", moudle.SysUserTokenKeyTime)
	// 存储用户id->token
	pipe.Set(moudle.Ctx, moudle.SysUserTokenKeyPrefix+id, token, moudle.SysUserTokenKeyTime)

	exec, err := pipe.Exec(moudle.Ctx)
	if err != nil {
		log.Println("Redis存储异常失败: %v", err, exec)
		c.JSON(500, Result.Err)
		return
	}
	// 返回登录成功, 将token返回
	c.JSON(200, Result.OK.WithData(gin.H{"token": token}))
	return
}
