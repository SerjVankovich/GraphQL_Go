package gql

import (
	"database/sql"
	"github.com/graphql-go/graphql"
)

func Schema(dataBase *sql.DB) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    QueryType(dataBase),
			Mutation: MutationType(dataBase),
		},
	)
}
