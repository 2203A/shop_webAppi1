package router

import (
	"SX1/shop_api/api/api"
	"github.com/gin-gonic/gin"
)

func Router(c *gin.Engine) {
	user := c.Group("user")
	{
		user.POST("/login", api.Login)
	}
}
