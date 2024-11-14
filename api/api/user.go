package api

import (
	"SX1/shop_api/api/cmd/global"
	"SX1/shop_api/api/initnal/formvalidata"
	userService "SX1/shop_api/api/initnal/service/user"
	"SX1/shop_api/api/initnal/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type User struct {
}

func Login(c *gin.Context) {
	//表单验证
	user := formvalidata.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	//调用service方法
	registration, err := userService.UserRegistration(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误1" + err.Error(),
		})
		return
	}
	//jwt生成token
	token, err := tools.GetJwtToken(global.AppConfig.Jwt.SecretKey, time.Now().Unix(), int64(global.AppConfig.Jwt.ExpireTime), strconv.FormatInt(registration.Id, 10))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误2" + err.Error(),
		})
		return
	}
	//返回
	c.JSON(http.StatusOK, gin.H{
		"msg":  "注册成功",
		"data": token,
	})
	return

}
