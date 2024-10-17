package api

import (
	"image-retrieval/internal/context"

	"github.com/gin-gonic/gin"
)

type ApiContext struct {
	*context.BaseContext
	GCtx *gin.Context
	Base
	ApiError error
}
type Base struct {
	Data any
}

func NewApiContext(ctx *gin.Context) *ApiContext {
	baseCtx := &context.BaseContext{}
	baseCtx.Init()
	return &ApiContext{
		Base:        Base{},
		BaseContext: baseCtx,
		GCtx:        ctx,
	}
}
