package user

import (
	c "context"
	"errors"
	"image-retrieval/internal/resource/database/model"
)

func (userContext *UserContext) GetUserInfo(_ c.Context) error {
	user, err := model.GetUser()
	userContext.user = user
	return err
}
func (userContext *UserContext) VerifyUser(_ c.Context) (err error) {
	// panic("not implemented")
	return errors.New("not implemented")
}
