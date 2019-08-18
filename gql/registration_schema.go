package gql

import (
	"github.com/graphql-go/graphql"
)

var RegistrationSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Mutation: RegistrationMutationType,
		Query:    RegistrationQueryType,
	},
)
