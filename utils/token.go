package utils

import (
	"os"
	"time"
	"yt-users-service/exception"
	"yt-users-service/helper"
	"yt-users-service/model/web"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(request web.TokenCreateRequest, value time.Duration) string {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))

	expiredTime := time.Now().Add(time.Minute * value)

	claims := &web.TokenClaims{
		UserId:    request.UserId,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtTokenSecret)
	helper.PanicError(err)

	return tokenStr
}

func ClaimsToken(userToken string) web.TokenClaims {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))

	claims := &web.TokenClaims{}

	token, err := jwt.ParseWithClaims(userToken, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtTokenSecret, nil
		},
	)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			panic(exception.NewUnauthorizedError(err.Error()))
		}
	}

	if !token.Valid {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	return *claims
}
