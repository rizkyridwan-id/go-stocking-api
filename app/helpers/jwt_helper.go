package helpers

import (
	"fmt"
	"stockingapi/app/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

type JWTHelper struct {
	secretKey []byte
	exp       time.Duration
}

func (j *JWTHelper) GenerateToken(userId string) string {
	claims := &Claims{
		UserName: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.exp)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err.Error()))
	}

	return tokenString
}

func (j *JWTHelper) ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, err
	}

	return claims, nil
}

func CreateJWTHelper(exp time.Duration) JWTHelper {
	configProps := configs.LoadConfig()
	return JWTHelper{
		secretKey: []byte(configProps.JWT_SECRET_KEY),
		exp:       exp,
	}
}
