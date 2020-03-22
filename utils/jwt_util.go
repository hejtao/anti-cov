package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	Key    = []byte("anti-cov-yidi")
	Issuer = "anti-cov"
)

// jwt
type JWTClaims struct {
	Id int
	jwt.StandardClaims
}

//生成管理员Token
func GenerateTokenString(id int) (string, error) {
	claims := JWTClaims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 3600*24*30,
			Issuer:    Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(Key)

	return tokenStr, err
}

// 解析Token
// token string -> token -> claims -> Id
func ParseTokenString(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Key, nil
	})

	if err != nil {
		return 0, errors.New("解析Token失败")
	}

	claims, ok := token.Claims.(*JWTClaims)

	if ok && token.Valid {
		return claims.Id, nil
	}

	return 0, errors.New("token无效")
}
