package main

import (
	"fmt"
	"log"

	"golang-todo-list/app/models"
	"golang-todo-list/config"
)

func main() {
	fmt.Println(models.Db)

	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")

	fmt.Println(models.Db)

	user01 := &models.User{}
	user01.Name = "test01"
	user01.Email = "test01@gmail.com"
	user01.PassWord = "test01"

	fmt.Println(user01)

	user01.CreateUser()

	user11, _ := models.GetUser(1)

	fmt.Println(user11)

	user11.Name = "updated_test"
	user11.Email = "updated_test@gmail.com"

	user11.UpdateUser()

	user11, _ = models.GetUser(1)

	fmt.Println(user11)

	user11.DeleteUser()

	user11, _ = models.GetUser(1)

	/* ToDo Listを作成するための処理開始 */
	user02 := &models.User{}
	user02.Name = "test02"
	user02.Email = "test02@gmail.com"
	user02.PassWord = "test02"

	fmt.Println(user02)

	user01.CreateUser()

	user22, _ := models.GetUser(2)
	user22.CreateTodo("First Todo")

}
