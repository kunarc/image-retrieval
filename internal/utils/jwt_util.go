package utils

import (
	"image-retrieval/internal/resource/database/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTUtil struct {
}

var SecretKey string

func init() {
	// cfg, err := ini.Load("./config/app.ini")
	// if err != nil {
	// 	fmt.Printf("Fail to read file: %v", err)
	// 	os.Exit(1)
	// }
	// SecretKey = cfg.Section("jwt").Key("secretKey").String()]
	SecretKey = "lkk"
}
func (JWTUtil) GenerateToken(user model.User) (string, error) {
	// 创建一个新的令牌对象，指定签名方法和声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间
	})

	// 使用你的 SecretKey 签名并获得完整的编码后的字符串令牌
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
