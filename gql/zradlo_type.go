package gql

import "github.com/graphql-go/graphql"

var ZradloType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Zradlo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
