package main

import (
	"./db"
	"./gql"
	"./handlers"
	"./utils"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	handlers2 "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

	r := mux.NewRouter()

	r.Handle("/zradlo", jwtMiddleware.Handler(handlers.GQLHandler(zradloSchema)))
	r.Handle("/register", handlers.GQLHandler(registerSchema(dataBase)))
	r.Handle("/login", handlers.GQLHandler(loginSchema))

	err = http.ListenAndServe(":8080", handlers2.LoggingHandler(os.Stdout, r))

	if err != nil {
		panic(err)
	}
}
