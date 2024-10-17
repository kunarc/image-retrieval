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
	User      *model.User
	LoginInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
}

func NewUserContext(ctx *gin.Context) *UserContext {
	return &UserContext{
		Base:       Base{User: &model.User{}},
		ApiContext: api.NewApiContext(ctx),
	}
}
