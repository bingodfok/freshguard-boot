package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type StandardClaims struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

type JwtAuth struct {
	SigningKey string
}

// GenToken 生成JWTToken
func (jwtAuth *JwtAuth) GenToken(claims StandardClaims) (string, error) {
	// 对称加密生成JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtAuth.SigningKey))
}

// ParseToken 解析JWTToken
func (jwtAuth *JwtAuth) ParseToken(tokenString string) (*StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtAuth.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*StandardClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
