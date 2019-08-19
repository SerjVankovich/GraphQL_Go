package gql

import (
	"database/sql"
	"github.com/graphql-go/graphql"
)

func RegistrationSchema(dataBase *sql.DB) graphql.Schema {
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Mutation: RegistrationMutationType(dataBase),
			Query:    RegistrationQueryType(dataBase),
		},
	)

	return schema
}
