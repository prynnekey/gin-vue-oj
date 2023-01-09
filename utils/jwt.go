package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var tokenKey = []byte("gin-gorm-oj")

// 生成token
func GenerateToken(identity, username string) (string, error) {
	userClaim := UserClaims{
		Identity:       identity,
		Username:       username,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &userClaim)
	tokenString, err := token.SignedString(tokenKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*UserClaims, error) {
	userClaim := UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &userClaim, func(t *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("analyse token error:%v", err)
	}

	return &userClaim, nil
}
