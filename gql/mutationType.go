package gql

import (
	"../models"
	"../utils"
	"errors"
	"github.com/graphql-go/graphql"
)

func createZradlo(zradla *[]*models.Zradlo) func(p graphql.ResolveParams) (i interface{}, e error) {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		name, nameOk := p.Args["name"].(string)
		price := p.Args["price"].(float64)

		id := utils.GetMaxID(*zradla) + 1
		zradlo := &models.Zradlo{}
		if nameOk {
			zradlo.Name = name
		}
		zradlo.Price = float32(price)
		zradlo.ID = int32(id)

		*zradla = append(*zradla, zradlo)

		return zradlo, nil
	}
}

func create(zradla *[]*models.Zradlo) *graphql.Field {
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
		Resolve: createZradlo(zradla),
	}
}

func update(zradla *[]*models.Zradlo) *graphql.Field {
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
		Resolve: updateZradla(zradla),
	}
}

func updateZradla(zradla *[]*models.Zradlo) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		id, idOk := p.Args["id"].(int)
		name, nameOk := p.Args["name"].(string)
		price, priceOk := p.Args["price"].(float64)

		if !idOk {
			return nil, errors.New(`cannot find "id" in query`)
		}
		zradlo, err := utils.FindByID(id, *zradla)

		if err != nil {
			return nil, err
		}

		if nameOk {
			zradlo.Name = name
		}

		if priceOk {
			zradlo.Price = float32(price)
		}

		return zradlo, nil
	}
}

func deleteQL(zradla *[]*models.Zradlo) *graphql.Field {
	return &graphql.Field{
		Type:        ZradloType,
		Description: "Delete one of zradla",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: deleteZradlo(zradla),
	}
}

func deleteZradlo(zradla *[]*models.Zradlo) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		id, idOk := p.Args["id"].(int)
		if !idOk {
			return nil, errors.New(`cannot find "id" in query`)
		}

		index, err := utils.FindID(id, *zradla)
		if err != nil {
			return nil, err
		}

		zr := *zradla
		zradlo := zr[index]
		*zradla = append(zr[:index], zr[index+1:]...)

		return zradlo, nil
	}
}

func MutationType(zradla *[]*models.Zradlo) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"create": create(zradla),
				"update": update(zradla),
				"delete": deleteQL(zradla),
			},
		},
	)
}
