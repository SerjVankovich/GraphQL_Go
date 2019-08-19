package db

import (
	"../models"
	"../utils"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"os"
)

var dberr = errors.New("you don't provide DataBase to function")
var zradloEmpty = errors.New("you don't provide zradlo to function")

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
	if db == nil {
		return nil, dberr
	}
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
	if db == nil {
		return nil, dberr
	}

	row := db.QueryRow(`SELECT * FROM public.zradlo WHERE "ID" = $1`, id)
	zradlo := new(models.Zradlo)
	err := row.Scan(&zradlo.ID, &zradlo.Price, &zradlo.Name)

	if err != nil {
		return nil, err
	}

	return zradlo, nil
}

func InsertZradlo(db *sql.DB, zradlo *models.Zradlo) (*models.Zradlo, error) {
	if db == nil {
		return nil, dberr
	}
	if zradlo == nil {
		return nil, zradloEmpty
	}
	_, err := db.Exec(`INSERT INTO public.zradlo("ID", "Price", "Name")VALUES (default, $2, $1)`,
		zradlo.Name,
		zradlo.Price)
	if err != nil {
		return nil, err
	}

	return zradlo, nil
}

func UpdateZradlo(db *sql.DB, zradlo *models.Zradlo) (*models.Zradlo, error) {
	if db == nil {
		return nil, dberr
	}
	if zradlo == nil {
		return nil, zradloEmpty
	}
	_, err := db.Exec(`UPDATE public.zradlo SET "Name" = $1, "Price" = $2 WHERE "ID" = $3`,
		zradlo.Name,
		zradlo.Price,
		zradlo.ID)

	if err != nil {
		return nil, err
	}

	return zradlo, nil
}

func DeleteZradlo(db *sql.DB, id int) error {
	if db == nil {
		return dberr
	}
	_, err := db.Exec(`DELETE FROM public.zradlo WHERE "ID" = $1`, id)
	return err
}

func InsertUser(db *sql.DB, user *models.User) error {
	if db == nil {
		return dberr
	}

	if user == nil {
		return errors.New("user field is empty")
	}

	_, err := db.Exec(`INSERT INTO public.users("id", "email", "password", "confirmed") VALUES (default, $1, $2, $3)`,
		user.Email,
		user.Password,
		false)

	return err
}

func UpdateUser(db *sql.DB, email string) (bool, error) {
	if db == nil {
		return false, dberr
	}

	_, err := db.Exec(`UPDATE public.users SET "confirmed" = $1 WHERE "email" = $2`, true, email)

	if err != nil {
		return false, err
	}

	return true, nil
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	if db == nil {
		return nil, dberr
	}
	row := db.QueryRow(`SELECT * FROM public.users WHERE "email" = $1`, email)

	user := new(models.User)

	err := row.Scan(&user.Id, &user.Email, &user.Confirmed, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func DeleteUser(db *sql.DB, email string) error {
	if db == nil {
		return dberr
	}

	_, err := db.Exec(`DELETE FROM public.users WHERE "email" = $1`, email)
	return err
}
