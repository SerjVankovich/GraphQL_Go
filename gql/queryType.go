package gql

import (
	"../db"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
)

func findById(dataBase *sql.DB) func(p graphql.ResolveParams) (i interface{}, e error) {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		id, ok := p.Args["id"].(int)
		if ok {
			return db.GetZradloByID(dataBase, id)
		}
		return nil, errors.New("zradlo not found")
	}
}

func getZradla(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(ZradloType),
		Description: "Get all zradla",
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return db.GetAllZradla(dataBase)
		},
	}
}

func getZradlo(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        ZradloType,
		Description: "Get zradlo by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: findById(dataBase),
	}
}

func QueryType(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"zradlo": getZradlo(dataBase),

				"zradla": getZradla(dataBase),
			},
		},
	)
}
