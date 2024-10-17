package middleware

import (
	"fmt"
	"image-retrieval/internal/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 创建一个 Gin 中间件，用于 JWT 验证
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 token
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}
		tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			// 确保 token 的签名方法是我们所期望的
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(utils.SecretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if claims, ok := tokenObj.Claims.(jwt.MapClaims); ok && tokenObj.Valid {
			// 如果 token 有效，并且你想在后续的处理程序中使用 token 中的信息
			// 你可以将信息添加到 Gin 的上下文中
			c.Set("userId", claims["userId"])

		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// 如果一切顺利，调用下一个中间件或处理程序
		c.Next()
	}
}
