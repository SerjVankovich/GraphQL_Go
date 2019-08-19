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
	var loginSchema = gql.LoginSchema(dataBase)

	http.HandleFunc("/zradlo", jwtMiddleware.Handler(handlers.GQLHandler(zradloSchema)).ServeHTTP)
	http.HandleFunc("/register", handlers.GQLHandler(registerSchema(dataBase)).ServeHTTP)
	http.HandleFunc("/login", handlers.GQLHandler(loginSchema).ServeHTTP)

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
