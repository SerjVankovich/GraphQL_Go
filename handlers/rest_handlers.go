package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func GetTokenHandler(secret []byte) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)

		claims["admin"] = true
		claims["name"] = "Sergey Vankovich"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		token.Claims = claims

		tokenString, _ := token.SignedString(secret)

		writer.Write([]byte(tokenString))
	})
}
