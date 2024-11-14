package routerGroup

import (
	"github.com/gin-gonic/gin"
	"shop-APi/intelnal/Api/userApi"
)

func RouterGroup(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		userGroup := v1.Group("/userGroup")
		{
			userGroup.POST("/register", userApi.UserRegister)
		}
	}
}
