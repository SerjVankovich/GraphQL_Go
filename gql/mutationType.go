package gql

import (
	"../db"
	"../models"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
	"strconv"
	"strings"
)

func createZradlo(dataBase *sql.DB) func(p graphql.ResolveParams) (i interface{}, e error) {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		name, nameOk := p.Args["name"].(string)
		price := p.Args["price"].(float64)

		zradlo := &models.Zradlo{}
		if nameOk {
			zradlo.Name = name
		}
		zradlo.Price = float32(price)

		return db.InsertZradlo(dataBase, zradlo)
	}
}

func create(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        ZradloType,
		Description: "Create new zradlo",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"price": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: createZradlo(dataBase),
	}
}

func update(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        ZradloType,
		Description: "Update one of zradla",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"price": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: updateZradla(dataBase),
	}
}

func updateZradla(dataBase *sql.DB) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		id, idOk := p.Args["id"].(int)
		name, nameOk := p.Args["name"].(string)
		price, priceOk := p.Args["price"].(float64)

		if !idOk {
			return nil, errors.New(`cannot find "id" in query`)
		}
		zradlo, err := db.GetZradloByID(dataBase, id)

		if err != nil {
			return nil, err
		}

		if nameOk {
			zradlo.Name = name
		}

		if priceOk {
			zradlo.Price = float32(price)
		}

		return db.UpdateZradlo(dataBase, zradlo)
	}
}

func deleteQL(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        ZradloType,
		Description: "Delete one of zradla",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: deleteZradlo(dataBase),
	}
}

func deleteZradlo(dataBase *sql.DB) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		id, idOk := p.Args["id"].(int)
		if !idOk {
			return nil, errors.New(`cannot find "id" in query`)
		}

		zradlo, err := db.GetZradloByID(dataBase, id)
		if err != nil {
			return nil, err
		}
		err = db.DeleteZradlo(dataBase, id)

		if err != nil {
			return nil, err
		}
		return zradlo, nil
	}
}

func deleteMore(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(ZradloType),
		Description: "Delete more than one zradlo",
		Args: graphql.FieldConfigArgument{
			"ids": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: deleteMoreZr(dataBase),
	}
}

func deleteMoreZr(dataBase *sql.DB) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		ids, idsOk := p.Args["ids"].(string)
		if !idsOk {
			return nil, errors.New(`cannot find "ids" in query`)
		}
		arrayIds := strings.Split(ids, ",")
		var deletedZradla []*models.Zradlo
		for _, id := range arrayIds {
			intId, err := strconv.Atoi(id)

			if err != nil {
				return nil, err
			}
			zradlo, err := db.GetZradloByID(dataBase, intId)

			if err != nil {
				return nil, err
			}

			deletedZradla = append(deletedZradla, zradlo)
			err = db.DeleteZradlo(dataBase, intId)
			if err != nil {
				return nil, err
			}
		}
		return deletedZradla, nil
	}
}

func MutationType(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"create":     create(dataBase),
				"update":     update(dataBase),
				"delete":     deleteQL(dataBase),
				"deleteMore": deleteMore(dataBase),
			},
		},
	)
}
