package main

import (
	"fmt"
	"log"

	"golang-todo-list/app/models"
	"golang-todo-list/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")

	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@gmail.com"
	u.PassWord = "test"

	fmt.Println(u)

	u.CreateUser()
}
