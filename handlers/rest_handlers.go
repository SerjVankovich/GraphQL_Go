package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
)

type AuthResponse struct {
	Content string `json:"content"`
	Error   string `json:"error"`
	Email   string `json:"email"`
}

func GoogleAuthHandler(config *oauth2.Config, oauthStateString string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		state, code := request.FormValue("state"), request.FormValue("code")
		content, err := getUserInfo(state, code, config, oauthStateString, "https://www.googleapis.com/oauth2/v2/userinfo?access_token=")
		content.Error = err.Error()
		_ = json.NewEncoder(writer).Encode(content)

	})
}
func GoogleLoginHandler(config *oauth2.Config, oauthStateString string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		url := config.AuthCodeURL(oauthStateString)
		http.Redirect(writer, request, url, http.StatusTemporaryRedirect)
	})
}

func getUserInfo(state string, code string, config *oauth2.Config, oauthStateString string, accessTokenUri string) (*AuthResponse, error) {
	if state != oauthStateString {
		return nil, errors.New("state_string isn't valid")
	}

	token, err := config.Exchange(context.Background(), code)

	if err != nil {
		return nil, fmt.Errorf("couldn't exchange token, err: %s", err.Error())
	}

	authResponse := &AuthResponse{Content: token.AccessToken, Email: token.Extra("email").(string), Error: ""}

	return authResponse, nil

}

func VkLoginHandler(oauthConfig *oauth2.Config, state string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		url := oauthConfig.AuthCodeURL(state)
		fmt.Println(url)
		http.Redirect(writer, request, url, http.StatusTemporaryRedirect)
	})
}

func VkAuthHandler(oauthConfig *oauth2.Config, oauthState string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		state, code := request.FormValue("state"), request.FormValue("code")
		email := request.FormValue("email")

		fmt.Println(email)

		content, err := getUserInfo(state, code, oauthConfig, oauthState, "")

		if err != nil {
			content.Error = err.Error()
		}
		_ = json.NewEncoder(writer).Encode(content)

	})
}
