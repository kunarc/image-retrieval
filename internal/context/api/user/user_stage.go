package user

import (
	c "context"
	"errors"
	"image-retrieval/internal/resource/database/model"
	"image-retrieval/internal/serialize"
	"image-retrieval/internal/utils"
)

func (userContext *UserContext) GetUserInfo(_ c.Context) error {
	if err := userContext.GCtx.ShouldBindJSON(&userContext.LoginInfo); err != nil {
		return err
	}
	user, err := model.GetUser(userContext.LoginInfo.Username)
	userContext.User = user
	return err
}
func (userContext *UserContext) VerifyUser(_ c.Context) error {
	if userContext.User == nil {
		err := errors.New("账号错误，请重新输入")
		return err
	}
	if userContext.User.Password != userContext.LoginInfo.Password {
		err := errors.New("密码错误，请重新输入")
		return err
	}
	jwtUtil := utils.JWTUtil{}
	token, err := jwtUtil.GenerateToken(*userContext.User)
	if err != nil {
		return err
	}
	userContext.Data = serialize.UserLogined{
		User:  *userContext.User,
		Token: token,
	}
	return nil
}
func (userContext *UserContext) RegisterUser(_ c.Context) (err error) {
	if err = userContext.GCtx.ShouldBindJSON(userContext.User); err != nil {
		return
	}
	err = model.CreateUser(userContext.User)
	return
}
