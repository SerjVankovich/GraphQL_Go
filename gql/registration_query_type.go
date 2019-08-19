package gql

import (
	"database/sql"
	"github.com/graphql-go/graphql"
)

func RegistrationQueryType(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"registers": &graphql.Field{
					Type:        graphql.NewList(RegistrationType),
					Description: "List of all registers",
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						//TODO Make realization
						return nil, nil
					},
				},
			},
		},
	)
}
