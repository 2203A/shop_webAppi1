package userApi

import (
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	global2 "shop-APi/cmd/global"
	"shop-APi/intelnal/formValidate/userReq"
	"shop-APi/intelnal/service/userSrv"
	"shop-APi/intelnal/tools/jwt"
	"time"
)

func UserRegister(c *gin.Context) {
	var r userReq.UserRegister
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	register, err := userSrv.UserRegister(r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "登录失败",
			"data":    err.Error(),
		})
		return
	}
	claims := jwt.CustomClaims{
		ID: uint(register.Id),
		StandardClaims: jwt2.StandardClaims{
			ExpiresAt: time.Now().Unix() + global2.InitConfig.JwtToken.AccessExpire,
		},
	}
	newJWT := jwt.NewJWT(global2.InitConfig.JwtToken.AccessSecret)
	token, err := newJWT.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "服务器内部错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "登录成功",
		"data":    token,
	})
}
