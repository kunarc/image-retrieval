package user

import (
	"image-retrieval/internal/context/api"
	"image-retrieval/internal/resource/database/model"

	"github.com/gin-gonic/gin"
)

type UserContext struct {
	*api.ApiContext
	Base
}

type Base struct {
	user *model.User
}

func NewUserContext(ctx *gin.Context) *UserContext {
	return &UserContext{
		Base:       Base{},
		ApiContext: api.NewApiContext(ctx),
	}
}
