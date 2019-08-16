package gql

import (
	"../models"
	"errors"
	"github.com/graphql-go/graphql"
)

func findById(zradla *[]*models.Zradlo) func(p graphql.ResolveParams) (i interface{}, e error) {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		id, ok := p.Args["id"].(int)
		if ok {
			for _, zradlo := range *zradla {
				if id == int(zradlo.ID) {
					return zradlo, nil
				}
			}
		}

		return nil, errors.New("zradlo not found")
	}
}

func getZradla(zradla *[]*models.Zradlo) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(ZradloType),
		Description: "Get all zradla",
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return zradla, nil
		},
	}
}

func getZradlo(zradla *[]*models.Zradlo) *graphql.Field {
	return &graphql.Field{
		Type:        ZradloType,
		Description: "Get zradlo by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: findById(zradla),
	}
}

func QueryType(zradla *[]*models.Zradlo) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"zradlo": getZradlo(zradla),

				"zradla": getZradla(zradla),
			},
		},
	)
}
