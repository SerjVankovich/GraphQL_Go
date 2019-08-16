package utils

import (
	"../models"
	"testing"
)

var zradla = []*models.Zradlo{
	{1, "Pizza", 600},
	{2, "Sushi", 1000},
	{3, "Borsch", 200}}

type testPairMaxID struct {
	zradla []*models.Zradlo
	value  interface{}
}

type testPairFindByID struct {
	zradla []*models.Zradlo
	value  *models.Zradlo
}

type testPairFindID struct {
	zradla []*models.Zradlo
	value  int
}

var testPairsMaxID = []testPairMaxID{
	{[]*models.Zradlo{}, 0},
	{zradla, 3},
}

var testPairsFindByID = []testPairFindByID{
	{[]*models.Zradlo{}, nil},
	{zradla, &models.Zradlo{ID: 2, Name: "Sushi", Price: 1000}},
}

var testPairsFindID = []testPairFindID{
	{[]*models.Zradlo{}, 0},
	{zradla, 1},
}

func TestGetMaxID(t *testing.T) {
	for _, pair := range testPairsMaxID {
		maxId := GetMaxID(pair.zradla)
		if maxId != pair.value {
			t.Error(
				"For", pair.zradla,
				"expected", pair.value,
				"got", maxId,
			)
		}
	}
}

func TestFindByID(t *testing.T) {
	for _, pair := range testPairsFindByID {
		zradlo, _ := FindByID(2, pair.zradla)
		if zradlo == nil && pair.value == nil {
			return
		}

		if zradlo.ID != pair.value.ID || zradlo.Name != pair.value.Name || zradlo.Price != pair.value.Price {
			t.Error(
				"For", pair.zradla,
				"expected", pair.value.ID, pair.value.Name, pair.value.Price,
				"got", zradlo.ID, zradlo.Name, zradlo.Price,
			)
		}
	}
}

func TestFindID(t *testing.T) {
	for _, pair := range testPairsFindID {
		index, _ := FindID(2, pair.zradla)
		if index != pair.value {
			t.Error(
				"For", pair.zradla,
				"expected", pair.value,
				"got", index,
			)
		}
	}
}
