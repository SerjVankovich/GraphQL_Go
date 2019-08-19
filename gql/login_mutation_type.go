package gql

import (
	"../db"
	"../models"
	"../utils"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
	"os"
)

func LoginMutationType(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"emPass": emPass(dataBase),
		},
	})
}

func emPass(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        LoginType,
		Description: "Login to service by email and password",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: emPassResolver(dataBase),
	}
}

func emPassResolver(dataBase *sql.DB) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		email, emailOk := p.Args["email"].(string)
		password, passwordOk := p.Args["password"].(string)
		if !emailOk {
			return nil, errors.New("email field not provided")
		}
		if !passwordOk {
			return nil, errors.New("password field not provided")
		}

		user, err := db.GetUserByEmail(dataBase, email)

		if err != nil {
			return nil, err
		}

		if user.Password != password {
			return nil, errors.New("invalid password")
		}

		path, err := os.Getwd()

		if err != nil {
			return nil, err
		}

		secretJson := utils.ParseSecret(path + "\\keys.json")

		token, err := utils.CreateToken(secretJson, email)

		return models.Login{token}, err
	}
}
