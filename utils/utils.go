package utils

import (
	"../models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
)

type config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Sslmode  string `json:"sslmode"`
}

type secret struct {
	JWTSecret string `json:"json-secret"`
}

type Sender struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetMaxID(zradla []*models.Zradlo) int {
	if len(zradla) == 0 {
		return 0
	}
	maxID := zradla[0].ID
	for _, zradlo := range zradla {
		if maxID < zradlo.ID {
			maxID = zradlo.ID
		}
	}

	return int(maxID)
}

func FindByID(id int, zradla []*models.Zradlo) (*models.Zradlo, error) {
	for _, zradlo := range zradla {
		if int(zradlo.ID) == id {
			return zradlo, nil
		}
	}

	return nil, errors.New(`cannot find zradlo on this "id"`)
}

func FindID(id int, zradla []*models.Zradlo) (int, error) {
	for key, zradlo := range zradla {
		if int(zradlo.ID) == id {
			return key, nil
		}
	}

	return 0, errors.New(`cannot find zradlo on this id`)
}

func DeleteByID(id int, zradla *[]*models.Zradlo) (*[]*models.Zradlo, *models.Zradlo) {
	zr := *zradla
	zradlo := zr[id]
	*zradla = append(zr[:id], zr[id+1:]...)

	return zradla, zradlo
}

func ParseConfig(path string) string {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(file)
	var config config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return "user=" +
		config.User +
		" password=" +
		config.Password +
		" dbname=" +
		config.Dbname +
		" sslmode=" +
		config.Sslmode
}

func ParseSecret(path string) []byte {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	var json_secret secret

	err = json.Unmarshal(byteValue, &json_secret)

	if err != nil {
		panic(err)
	}

	return []byte(json_secret.JWTSecret)
}

func SendEmail(email string, password string) {
	path, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	file, err := os.Open(path + "\\keys.json")

	if err != nil {
		panic(err)
	}

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	var sender Sender

	err = json.Unmarshal(byteValue, &sender)

	if err != nil {
		panic(err)
	}

	auth := smtp.PlainAuth("", sender.Email, sender.Password, "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, sender.Email, []string{email}, []byte("Привет"))

	fmt.Println(err.Error())

}
