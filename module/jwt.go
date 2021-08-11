package module

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const (
	SECRETKEY         = "123123"
	MAXAGE            = 3600 * 24
	CACHE_BLACK_TOKEN = "black.token."
)

type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

// 生成Token
func (c *CustomClaims) GetToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(SECRETKEY))
}

//解析token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
