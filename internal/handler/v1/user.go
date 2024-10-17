package v1

import (
	"image-retrieval/internal/context/api/user"

	"github.com/gin-gonic/gin"
)

func LoginVerify(ctx *gin.Context) {
	context := user.NewUserContext(ctx)
	context.AddBaseHandler(context.GetUserInfo, "getUserInfo").AddBaseHandler(context.VerifyUser, "verifyUser")
	context.AddDeferHandler(context.SendReponse, "sendResponse")
	context.Run()
}
func Register(ctx *gin.Context) {
	context := user.NewUserContext(ctx)
	context.AddBaseHandler(context.RegisterUser, "registerUser")
	context.AddDeferHandler(context.SendReponse, "sendResponse")
	context.Run()
}
