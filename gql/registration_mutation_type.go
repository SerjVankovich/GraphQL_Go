package gql

import (
	"../utils"
	"errors"
	"github.com/graphql-go/graphql"
)

func registerResolver(p graphql.ResolveParams) (i interface{}, e error) {
	email, emailOk := p.Args["email"].(string)
	password, passwordOk := p.Args["password"].(string)
	if !emailOk {
		return nil, errors.New("not provided email in query")
	}
	if !passwordOk {
		return nil, errors.New("not provided password in query")
	}

	utils.SendEmail(email, password)

	return nil, nil

}

var register = &graphql.Field{
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
	Resolve: registerResolver,
}

var RegistrationMutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"register": register,
		},
	},
)
