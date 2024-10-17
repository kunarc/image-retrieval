package serialize

import "image-retrieval/internal/resource/database/model"

type SearchResult struct {
	Records []model.Image `json:"records"`
	Total   int64         `json:"total"`
}
