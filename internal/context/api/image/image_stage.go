package image

import (
	"context"
	c "context"
	"encoding/json"
	"fmt"
	"image-retrieval/internal/config"
	"image-retrieval/internal/resource/database"
	"image-retrieval/internal/resource/database/model"
	"image-retrieval/internal/resource/es"
	"image-retrieval/internal/serialize"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/olivere/elastic/v7"
)

func (imageContext *ImageContext) GetRecommendImage(_ c.Context) error {
	images, total, err := model.GetImageByCount(8)
	if err != nil {
		return err
	}
	imageContext.Data = serialize.SearchResult{
		Records: images,
		Total:   total,
	}
	return nil
}

func (imageContext *ImageContext) GetRecentImage(_ c.Context) error {
	images, total, err := model.GetImageByCount(8)
	if err != nil {
		return err
	}
	imageContext.Data = serialize.SearchResult{
		Records: images,
		Total:   total,
	}
	return nil
}

func (imageContext *ImageContext) BuildQuery(_ c.Context) error {
	if err := imageContext.GCtx.ShouldBind(imageContext.SearchInfo); err != nil {
		return err
	}
	var (
		specificType = make([]string, 0)
		text         = imageContext.SearchInfo.Text
		pictureType  = imageContext.SearchInfo.Type
		figures      = imageContext.SearchInfo.FilterOption.Figures
		meanings     = imageContext.SearchInfo.FilterOption.Meanings
		years        = imageContext.SearchInfo.FilterOption.Years
		museums      = imageContext.SearchInfo.FilterOption.Meanings
	)
	query := elastic.NewBoolQuery()
	mustQueryList := make([]elastic.Query, 0)
	if text != "" {
		fileds := []string{"name", "artist", "details"}
		mustQueryList = append(mustQueryList, elastic.NewMultiMatchQuery(text, fileds...))
	}
	if pictureType != "" {
		mustQueryList = append(mustQueryList, elastic.NewTermQuery("type", pictureType))
	}
	if len(years) > 0 {
		boolQuery := elastic.NewBoolQuery()
		termQueryList := make([]elastic.Query, 0)
		for _, y := range years {
			termQueryList = append(termQueryList, elastic.NewTermQuery("year", y))
		}
		mustQueryList = append(mustQueryList, boolQuery.Should(termQueryList...))
	}
	specificType = append(specificType, figures...)
	specificType = append(specificType, meanings...)
	if len(specificType) > 0 {
		boolQuery := elastic.NewBoolQuery()
		termQueryList := make([]elastic.Query, 0)
		for _, sp := range specificType {
			termQueryList = append(termQueryList, elastic.NewTermQuery("specificType", sp))
		}
		mustQueryList = append(mustQueryList, boolQuery.Should(termQueryList...))
	}
	if len(museums) > 0 {
		boolQuery := elastic.NewBoolQuery()
		termQueryList := make([]elastic.Query, 0)
		for _, mu := range museums {
			termQueryList = append(termQueryList, elastic.NewTermQuery("museum", mu))
		}
		mustQueryList = append(mustQueryList, boolQuery.Should(termQueryList...))
	}
	query.Must(mustQueryList...)
	imageContext.Query = query
	return nil
}
func (imageContext *ImageContext) SearchImageByFilter(_ c.Context) error {
	page := imageContext.SearchInfo.Page
	pageSize := imageContext.SearchInfo.PageSize
	yearsOrder := imageContext.SearchInfo.YearsOrder
	matchOrder := imageContext.SearchInfo.MatchOrder
	service := es.ESClient.Search().
		Index("images").
		From((page - 1) * pageSize).
		Query(imageContext.Query)
	if yearsOrder != 0 {
		if yearsOrder == 1 {
			service.Sort("year", true)
		} else {
			service.Sort("year", false)
		}
	}
	if matchOrder != 0 {
		if matchOrder == 1 {
			service.Sort("_score", true)
		} else {
			service.Sort("_score", false)
		}
	}
	res, err := service.Do(c.Background())
	if err != nil {
		return err
	}
	searchRes := new(serialize.SearchResult)
	records := make([]model.Image, 0)
	for _, hit := range res.Hits.Hits {
		var record model.Image
		src, err := json.Marshal(hit.Source)
		if err != nil {
			return err
		}
		err = json.Unmarshal(src, &record)
		if err != nil {
			return nil
		}
		records = append(records, record)
	}
	searchRes.Total = res.Hits.TotalHits.Value
	searchRes.Records = records
	imageContext.Data = searchRes
	return nil
}

func (imageContext *ImageContext) BindImage(_ c.Context) (err error) {
	var form *multipart.Form
	form, err = imageContext.GCtx.MultipartForm()
	value := form.Value
	u, _ := strconv.Atoi(value["id"][0])
	imageContext.Image.ID = uint(u)
	imageContext.Image.Name = value["name"][0]
	imageContext.Image.Artist = value["artist"][0]
	imageContext.Image.Year = value["year"][0]
	imageContext.Image.Type = value["type"][0]
	imageContext.Image.SpecificType = value["specificType"][0]
	imageContext.Image.Size = value["size"][0]
	imageContext.Image.Museum = value["museum"][0]
	imageContext.Image.Details = value["details"][0]
	parsedUint64, _ := strconv.ParseUint(value["view"][0], 10, 32)
	imageContext.Image.View = uint32(parsedUint64)
	parsedUint64, _ = strconv.ParseUint(value["downLoad"][0], 10, 32)

	imageContext.Image.Download = uint32(parsedUint64)
	return
}

func (imageContext *ImageContext) BindAndSaveFiles(_ c.Context) (err error) {
	var form *multipart.Form
	form, err = imageContext.GCtx.MultipartForm()
	files := form.File["files"]
	pictrue := make([]string, 3)
	for i, file := range files {
		id := uuid.New().String()
		fileName := file.Filename
		etra := strings.Split(fileName, ".")[1]
		path := fmt.Sprintf("./static/images/%s.%s", id, etra)
		imageContext.GCtx.SaveUploadedFile(file, path)
		pictrue[i] = fmt.Sprintf("%s:%s/images/%s.%s", config.Config.GetString("service.host"), config.Config.GetString("service.port"), fileName, etra)
	}
	imageContext.Image.Picture = pictrue[0]
	imageContext.Image.Rubbing = pictrue[1]
	imageContext.Image.Picture3D = pictrue[2]
	return
}

func (imageContext *ImageContext) InsertToDb(_ c.Context) (err error) {
	err = database.DB.Omit("created_at").Save(imageContext.Image).Error
	return
}

func (imageContext *ImageContext) InsertToEs(_ c.Context) (err error) {
	var resp *elastic.IndexResponse
	resp, err = es.ESClient.Index().Index(imageContext.Image.Index()).Id(strconv.Itoa(int(imageContext.Image.ID))).BodyJson(imageContext.Image).Do(context.Background())
	if err != nil {
		return
	}
	if resp.Result == "created" {
		println("insert successful")
	} else {
		println("update successful")
	}
	return
}
