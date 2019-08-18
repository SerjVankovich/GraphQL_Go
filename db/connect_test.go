package db

import (
	"../utils"
	"database/sql"
	"os"
	"testing"
)

func connect() *sql.DB {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	connectString := utils.ParseConfig(path + "\\config.json")

	db, err := sql.Open("postgres", connectString)

	if err != nil {
		panic(err)
	}
	return db
}

var db = connect()

func TestGetAllZradla(t *testing.T) {
	_, err := GetAllZradla(nil)

	if err.Error() != dberr.Error() {
		t.Error(
			"For", nil,
			"expected", dberr,
			"got", err.Error(),
		)
	}
}

func TestGetZradloByID(t *testing.T) {
	_, err := GetZradloByID(nil, 0)

	if err.Error() != dberr.Error() {
		t.Error(
			"For", nil,
			"expected", dberr,
			"got", err.Error(),
		)
	}

	_, err = GetZradloByID(db, 1)

	errStr := "sql: no rows in result set"

	if err.Error() != errStr {
		t.Error(
			"For row with id:", 1,
			"expected", errStr,
			"got", err.Error(),
		)
	}

	zradlo, err := GetZradloByID(db, 2)
	if zradlo.Name != "Udon" || zradlo.Price != 400 {
		t.Error(
			"For", 2,
			"expected", "Udon", "400",
			"got", zradlo.Name, zradlo.Price,
		)
	}
}

func TestInsertZradlo(t *testing.T) {
	_, err := InsertZradlo(nil, nil)

	if err.Error() != dberr.Error() {
		t.Error(
			"For", nil,
			"expected", dberr.Error(),
			"got", err.Error(),
		)
	}

	_, err = InsertZradlo(db, nil)
	if err.Error() != zradloEmpty.Error() {
		t.Error(
			"For", db, nil,
			"expected", zradloEmpty.Error(),
			"got", err.Error(),
		)
	}
}
func TestUpdateZradlo(t *testing.T) {
	_, err := UpdateZradlo(nil, nil)

	if err.Error() != dberr.Error() {
		t.Error(
			"For", nil,
			"expected", dberr.Error(),
			"got", err.Error(),
		)
	}

	_, err = UpdateZradlo(db, nil)
	if err.Error() != zradloEmpty.Error() {
		t.Error(
			"For", db, nil,
			"expected", zradloEmpty.Error(),
			"got", err.Error(),
		)
	}
}

func TestDeleteZradlo(t *testing.T) {
	err := DeleteZradlo(nil, 0)

	if err.Error() != dberr.Error() {
		t.Error(
			"For", nil,
			"expected", dberr.Error(),
			"got", err.Error(),
		)
	}
}
