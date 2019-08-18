package gql

import (
	"database/sql"
	"github.com/graphql-go/graphql"
)

func ZradloSchema(dataBase *sql.DB) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    ZradloQueryType(dataBase),
			Mutation: ZradloMutationType(dataBase),
		},
	)
}
