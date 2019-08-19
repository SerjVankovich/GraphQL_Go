package gql

import "github.com/graphql-go/graphql"

var RegistrationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Registartion",
	Fields: graphql.Fields{
		"regLink": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var CompleteRegistrationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CompleteRegistration",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.String,
		},
	},
})
