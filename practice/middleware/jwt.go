package middleware

import (
	"time"

	"github.com/DesistDaydream/GoGin/practice/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

// 签名所需的私钥。如果想要解析已签名的 Token，则私钥必须相同，否则解析将会失败
var mySigningKey = []byte("AllYourBase")

type CustomClaims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// GenerateToken 生成 Token
func GenerateToken(userInfo *database.User) (string, error) {
	// Create the Claims
	claims := CustomClaims{
		userInfo.Name,
		jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Id:        "",
			IssuedAt:  0,
			Issuer:    "desistdaydream",
			NotBefore: 0,
			Subject:   "",
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := jwtToken.SignedString(mySigningKey)
	logrus.Debugf("%v 用户生成的 Token 为 %v\n", userInfo.Name, token)
	return token, nil
}

// ParseToken 解析 Token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
