package gql

import (
	"../models"
	"github.com/graphql-go/graphql"
)

func Schema(zradla *[]*models.Zradlo) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    QueryType(zradla),
			Mutation: MutationType(zradla),
		},
	)
}
