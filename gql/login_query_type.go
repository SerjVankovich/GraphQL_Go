package gql

import (
	"github.com/graphql-go/graphql"
)

func LoginQueryType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"login": &graphql.Field{
				Type:        LoginType,
				Description: "mock func",
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					return nil, nil
				},
			},
		},
	})
}
