package utils

import (
	"../models"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/mail"
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

type hmac_secret struct {
	HMACSecret string `json:"hmac-secret"`
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

func ParseHMAC(path string) ([]byte, error) {
	if path == "" {
		return nil, errors.New("you don't provide path to function")
	}
	file, err := os.Open(path + "\\keys.json")

	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	var hmacSecret hmac_secret

	err = json.Unmarshal(byteValue, &hmacSecret)

	if err != nil {
		return nil, err
	}

	return []byte(hmacSecret.HMACSecret), nil
}

func SendEmail(email string, data string) error {
	path, err := os.Getwd()

	if err != nil {
		return err
	}

	file, err := os.Open(path + "\\keys.json")

	if err != nil {
		return err
	}

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	var sender Sender

	err = json.Unmarshal(byteValue, &sender)

	if err != nil {
		return err
	}

	from := mail.Address{"", sender.Email}
	to := mail.Address{"", email}
	subj := "Registration"
	body := "To complete registration go to \n http://localhost:8080/register?query=mutation+_{completeRegister(hash:\"" +
		data + "\"){successful}}"

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := "smtp.gmail.com:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", sender.Email, sender.Password, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()

	return nil

}

func Encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func Decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func createHash(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))

	return hex.EncodeToString(hash.Sum(nil))
}
