package main

import (
	"./db"
	"./gql"
	"./handlers"
	"./utils"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
)

func main() {
	dataBase := db.Connect()
	defer dataBase.Close()

	path, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	var secret = utils.ParseSecret(path + "\\keys.json")

	var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (i interface{}, e error) {
			return secret, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	var zradloSchema, _ = gql.ZradloSchema(dataBase)
	var registerSchema = gql.RegistrationSchema

	http.HandleFunc("/zradlo", jwtMiddleware.Handler(handlers.GQLHandler(zradloSchema)).ServeHTTP)
	http.HandleFunc("/get-token", handlers.GetTokenHandler(secret).ServeHTTP)

	//register doesn't work
	http.HandleFunc("/register", handlers.GQLHandler(registerSchema).ServeHTTP)

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
