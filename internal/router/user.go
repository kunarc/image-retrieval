package router

import (
	v1 "image-retrieval/internal/handler/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/login", v1.LoginVerify)
		user.POST("/register", v1.Register)
	}
}
