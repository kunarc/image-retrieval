package image

import (
	"image-retrieval/internal/context/api"
	"image-retrieval/internal/resource/database/model"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type ImageContext struct {
	*api.ApiContext
	Base
}
type Base struct {
	Image      *model.Image
	SearchInfo *SearchInfo
	Query      elastic.Query
	BindData   BindData
	Files      []multipart.File
}
type Filter struct {
	Figures  []string // 图像砖细分类型
	Meanings []string // 文字砖细分类型
	Years    []string // 年代
	Museum   []string // 博物馆
}
type SearchInfo struct {
	FilterOption Filter
	Page         int
	PageSize     int
	MatchOrder   int
	YearsOrder   int
	Text         string
	Type         string
}
type BindData struct {
}

func NewImageContext(ctx *gin.Context) *ImageContext {
	return &ImageContext{
		Base:       Base{Image: &model.Image{}, SearchInfo: &SearchInfo{}},
		ApiContext: api.NewApiContext(ctx),
	}
}
