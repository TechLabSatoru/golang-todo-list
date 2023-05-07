package models

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID int
	UUID string
	Name string
	Email string
	PassWord string
	CreateAt time.Time
}

func (u *User) CreateUser() (err error) {

	cmd := fmt.Sprintf(`INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at) VALUES ($1, $2, $3, $4, $5)`)

	result, err := Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.PassWord), time.Now().Format("2006-01-02 15:04:05.999999-07:00"))
	if err != nil {
		log.Fatalln(err)
	}
	if result != nil {
		numRows, _ := result.RowsAffected()
		fmt.Printf("%d rows affected\n", numRows)
	}
	return err
}
