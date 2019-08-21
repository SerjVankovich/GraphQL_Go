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
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/vk"
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

	var oauthStateString = utils.GetRandomString(20)

	var secret = utils.ParseSecret(path + "\\keys.json")
	var googleCredentials = utils.ParseGoogleCredentials(path + "\\keys.json")
	var vkCredentials = utils.ParseVkCredentials(path + "\\keys.json")

	var googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/googleAuth",
		ClientID:     googleCredentials.GoogleClientId,
		ClientSecret: googleCredentials.GoogleSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	var vkOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/vkAuth",
		ClientID:     vkCredentials.ClientId,
		ClientSecret: vkCredentials.ClientSecret,
		Scopes:       []string{"email"},
		Endpoint:     vk.Endpoint,
	}

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

	r.Handle("/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		template := `
			<h1>
				<a href="/loginGoogle">Login To google</a>
				<a href="/loginVk">Login to VK</a>
			</h1>
					`

		writer.Write([]byte(template))
	}))
	r.Handle("/zradlo", jwtMiddleware.Handler(handlers.GQLHandler(zradloSchema)))
	r.Handle("/register", handlers.GQLHandler(registerSchema(dataBase)))
	r.Handle("/login", handlers.GQLHandler(loginSchema))

	// Google oauth 2.0 need to make it works
	r.Handle("/loginGoogle", handlers.GoogleLoginHandler(googleOauthConfig, oauthStateString))
	r.Handle("/googleAuth", handlers.GoogleAuthHandler(googleOauthConfig, oauthStateString))

	// VK oauth 2.0
	r.Handle("/loginVk", handlers.VkLoginHandler(vkOauthConfig, oauthStateString))
	r.Handle("/vkAuth", handlers.VkAuthHandler(vkOauthConfig, oauthStateString))

	err = http.ListenAndServe(":8080", handlers2.LoggingHandler(os.Stdout, r))

	if err != nil {
		panic(err)
	}
}
