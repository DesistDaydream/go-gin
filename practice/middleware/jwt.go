package middleware

import (
	"github.com/DesistDaydream/GoGin/practice/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type CustomClaims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// GenerateToken 生成 JWT
func GenerateToken(userInfo *database.User) (string, error) {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := CustomClaims{
		userInfo.Name,
		jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: 15000,
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
// func ParseToken(token string) (*jwt.StandardClaims, error) {
// 	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
// 		return []byte(config.Secret), nil
// 	})
// 	if err == nil && jwtToken != nil {
// 		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
// 			return claim, nil
// 		}
// 	}
// 	return nil, err
// }
