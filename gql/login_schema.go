package gql

import (
	"database/sql"
	"github.com/graphql-go/graphql"
)

func LoginSchema(dataBase *sql.DB) graphql.Schema {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    LoginQueryType(),
		Mutation: LoginMutationType(dataBase),
	})

	return schema
}
