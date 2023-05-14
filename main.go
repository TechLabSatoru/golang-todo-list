package main

import (
	"fmt"
	// "log"

	"golang-todo-list/app/models"
	"golang-todo-list/app/controllers"
	// "golang-todo-list/config"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
	/*
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

	user02 := &models.User{}
	user02.Name = "test02"
	user02.Email = "test02@gmail.com"
	user02.PassWord = "test02"

	fmt.Println(user02)

	user02.CreateUser()

	user22, _ := models.GetUser(2)
	user22.CreateTodo("First Todo")

	todo01, _ := models.GetTodo(1)
	
	user22, _ := models.GetUser(2)
	user22.CreateTodo("Second Todo")

	todos, _ := models.GetTodos()
	for _, v := range todos {
		fmt.Println(v)
	}
	
	user03 := &models.User{}
	user03.Name = "test03"
	user03.Email = "test03@gmail.com"
	user03.PassWord = "test03"

	fmt.Println(user03)

	user03.CreateUser()

	user23, _ := models.GetUser(3)
	user23.CreateTodo("Third Todo")

	user22, _ := models.GetUser(2)
	todos, _ := user22.GetTodosByUser()
	for _, v := range todos {
		fmt.Println(v)
	}

	todo01, _ := models.GetTodo(1)
	todo01.Content = "Update Todo"
	todo01.UpdateTodo()

	todo01.DeleteTodo()
	*/
}
