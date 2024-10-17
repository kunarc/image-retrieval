package v1

import (
	"image-retrieval/internal/context/api/image"

	"github.com/gin-gonic/gin"
)

func RecommendImage(c *gin.Context) {
	context := image.NewImageContext(c)
	context.AddDeferHandler(context.SendReponse, "sendResponse")
	context.AddDeferHandler(context.GetRecommendImage, "getRecommendImage")
	context.Run()
}
func RecentImage(c *gin.Context) {
	context := image.NewImageContext(c)
	context.AddBaseHandler(context.GetRecentImage, "getRecentImage")
	context.AddDeferHandler(context.SendReponse, "sendResponse")
	context.Run()
}
func SearchImage(c *gin.Context) {
	context := image.NewImageContext(c)
	context.AddBaseHandler(context.BuildQuery, "buildQuery").AddBaseHandler(context.SearchImageByFilter, "searchImageByFilter")
	context.AddDeferHandler(context.SendReponse, "sendResponse")
	context.Run()
}
func InsertImage(c *gin.Context) {
	context := image.NewImageContext(c)
	context.AddDeferHandler(context.SendReponse, "sendResponse")
	context.AddBaseHandler(context.BindImage, "bindImage").AddBaseHandler(context.BindAndSaveFiles, "bindAndSaveFiles").
		AddBaseHandler(context.InsertToDb, "insertToDb").AddBaseHandler(context.InsertToEs, "insertToEs")
	context.Run()

}
