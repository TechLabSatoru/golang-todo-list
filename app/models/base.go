package models

import (
	"fmt"
	"log"

	"crypto/sha1"
	"database/sql"

	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
)

func init() {
	Db, err = sql.Open("postgres", "host=docker.for.mac.host.internal user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		uuid VARCHAR(255) NOT NULL UNIQUE,
		name VARCHAR(255),
		email VARCHAR(255),
		password VARCHAR(255),
		created_at TIMESTAMP);`, tableNameUser)

	result, err := Db.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}
	if result != nil {
		numRows, _ := result.RowsAffected()
		fmt.Printf("%d rows affected\n", numRows)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		content TEXT,
		user_id INTEGER,
		created_at TIMESTAMP);`, tableNameTodo)
	
	Db.Exec(cmdT)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
