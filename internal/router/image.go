package router

import (
	v1 "image-retrieval/internal/handler/v1"

	"github.com/gin-gonic/gin"
)

func InitImageRouter(r *gin.Engine) {
	image := r.Group("/search")
	// image.Use(middleware.JWTAuthMiddleware())
	{
		image.GET("/recom", v1.RecommendImage)
		image.GET("/recent", v1.RecentImage)
		image.POST("/search", v1.SearchImage)
		image.POST("/insert", v1.InsertImage)
	}
}
