package models

import (
	"time"
	"log"
)

type Todo struct {
	ID int
	Content string
	UserID int
	CreateAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `INSERT INTO todos (content, user_id, created_at) VALUES($1, $2, $3)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}

	cmd := `SELECT id, content, user_id, created_at FROM todos WHERE id = $1`

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID, 
		&todo.Content, 
		&todo.UserID, 
		&todo.CreateAt)

	if err != nil {
		log.Fatalln(err)
	}

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos`

	rows, err := Db.Query(cmd)

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID, 
			&todo.Content,
			&todo.UserID,
			&todo.CreateAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	
	rows.Close()

	return todos, err
}
