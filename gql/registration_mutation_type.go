package gql

import (
	"../db"
	"../models"
	"../utils"
	"database/sql"
	"encoding/hex"
	"errors"
	"github.com/graphql-go/graphql"
	"os"
	"strings"
)

func registerResolver(dataBase *sql.DB) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		email, emailOk := p.Args["email"].(string)
		password, passwordOk := p.Args["password"].(string)
		if !emailOk {
			return nil, errors.New("not provided email in query")
		}
		if !passwordOk {
			return nil, errors.New("not provided password in query")
		}

		user, err := db.GetUserByEmail(dataBase, email)

		if user != nil {
			return nil, errors.New("user with this email exists")
		}

		if err != nil {
			return nil, err
		}

		path, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		hmacSecret, err := utils.ParseHMAC(path)

		hash := utils.Encrypt([]byte(email+" "+password), string(hmacSecret))

		hasStr := hex.EncodeToString(hash)

		reg := models.Registration{RegLink: hasStr}

		err = utils.SendEmail(email, hasStr)

		if err != nil {
			return nil, err
		}

		user = &models.User{Email: email, Password: password}

		err = db.InsertUser(dataBase, user)

		if err != nil {
			return nil, err
		}

		return reg, nil

	}
}

func register(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        RegistrationType,
		Description: "Register to new user",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: registerResolver(dataBase),
	}
}

func RegistrationMutationType(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"register":         register(dataBase),
				"completeRegister": completeRegister(dataBase),
			},
		},
	)
}

func completeRegister(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        CompleteRegistartionType,
		Description: "Complete Register of new User",
		Args: graphql.FieldConfigArgument{
			"hash": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: completeRegisterResolver(dataBase),
	}
}

func completeRegisterResolver(dataBase *sql.DB) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		path, err := os.Getwd()

		if err != nil {
			return nil, err
		}

		hash, hashOk := p.Args["hash"].(string)

		if !hashOk {
			return nil, errors.New("hash field not provided")
		}

		byteHash, err := hex.DecodeString(hash)

		if err != nil {
			return nil, err
		}

		hmac_secret, err := utils.ParseHMAC(path)

		if err != nil {
			return nil, err
		}

		emailPass := string(utils.Decrypt(byteHash, string(hmac_secret)))

		email := strings.Split(emailPass, " ")[0]

		ok, _ := db.UpdateUser(dataBase, email)

		if !ok {
			err := db.DeleteUser(dataBase, email)
			if err != nil {
				return nil, err
			}
		}

		completeReg := models.CompleteRegistration{Successful: ok}

		return completeReg, nil

	}
}
