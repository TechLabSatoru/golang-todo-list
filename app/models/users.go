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

func GetUser(id int) (user User, err error) {
	user = User{}

	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE id = $1`

	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreateAt,
	)

	return user, err
}


func (u *User) UpdateUser() (err error) {
	cmd := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `DELETE FROM users WHERE id = $1`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
