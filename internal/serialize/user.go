package serialize

import "image-retrieval/internal/resource/database/model"

type UserLogined struct {
	model.User
	Token string
}
