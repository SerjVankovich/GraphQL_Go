package db

import (
	"../models"
	"../utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func Connect() *sql.DB {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	connectString := utils.ParseConfig(path + "\\db\\config.json")

	db, err := sql.Open("postgres", connectString)

	if err != nil {
		panic(err)
	}
	return db
}

func GetAllZradla(db *sql.DB) (*[]*models.Zradlo, error) {
	var zradla []*models.Zradlo

	rows, err := db.Query("SELECT * FROM public.zradlo")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		zradlo := new(models.Zradlo)
		err = rows.Scan(&zradlo.ID, &zradlo.Price, &zradlo.Name)

		if err != nil {
			return nil, err
		}

		zradla = append(zradla, zradlo)
	}

	return &zradla, nil
}

func GetZradloByID(db *sql.DB, id int) (*models.Zradlo, error) {
	row := db.QueryRow(`SELECT * FROM public.zradlo WHERE "ID" = $1`, id)
	zradlo := new(models.Zradlo)
	err := row.Scan(&zradlo.ID, &zradlo.Price, &zradlo.Name)

	if err != nil {
		return nil, err
	}

	return zradlo, nil
}

func InsertZradlo(db *sql.DB, zradlo *models.Zradlo) (*models.Zradlo, error) {
	_, err := db.Exec(`INSERT INTO public.zradlo("ID", "Price", "Name")VALUES (default, $2, $1)`,
		zradlo.Name,
		zradlo.Price)
	if err != nil {
		return nil, err
	}

	return zradlo, nil
}

func UpdateZradlo(db *sql.DB, zradlo *models.Zradlo) (*models.Zradlo, error) {
	_, err := db.Exec(`UPDATE public.zradlo SET "Name" = $1, "Price" = $2 WHERE "ID" = $3`,
		zradlo.Name,
		zradlo.Price,
		zradlo.ID)

	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return zradlo, nil
}

func DeleteZradlo(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM public.zradlo WHERE "ID" = $1`, id)
	return err
}
