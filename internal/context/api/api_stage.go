package api

import (
	c "context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctx *ApiContext) SendReponse(_ c.Context) (err error) {
	if ctx.BaseError != nil {
		ctx.GCtx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  ctx.BaseError.Error(),
		})
		return
	}
	ctx.GCtx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": ctx.Data,
	})
	return
}
