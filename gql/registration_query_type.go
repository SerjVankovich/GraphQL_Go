package gql

import "github.com/graphql-go/graphql"

var RegistrationQueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"registers": &graphql.Field{
				Type:        graphql.NewList(RegistrationType),
				Description: "List of all registers",
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					return nil, nil
				},
			},
		},
	},
)
