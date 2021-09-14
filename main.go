package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		log.Panic("You got a err")
	}
	setup(db)

	flag := os.Args[1]
	argument := os.Args[2:]
	argStr := strings.Join(argument, " ")

	if flag == "add" && argument != nil {
		todo := Todos{Name: argStr}
		newTodo(db, &todo)
	}

	if flag == "ls" {
		listTodo(db)
	}

	if flag == "delall" {
		deleteAll(db)
	}

	if flag == "del" && argument != nil {
		deleteTodo(argStr, db)
	}
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&Todos{})
}

func newTodo(db *gorm.DB, todo *Todos) {
	db.Create(todo)
	fmt.Printf("New Todo is created with Id %d\n", todo.ID)
}

func listTodo(db *gorm.DB) {
	// Remeber : always pass a struct to be filled.
	var todos []Todos
	db.Find(&todos)
	for i, y := range todos {
		fmt.Printf("%d. %s - %d \n", i+1, y.Name, y.ID)
	}
}

func deleteTodo(key string, db *gorm.DB) {
	db.Unscoped().Delete(&Todos{}, key)
	fmt.Printf("Deleted todos with ID %s Succesfully \n", key)
}

func deleteAll(db *gorm.DB) {
	var todos []Todos
	db.Find(&todos)
	for _, y := range todos {
		db.Unscoped().Delete(&Todos{}, y.ID)
	}
	fmt.Println("Deleted All todos Succesfully")
}
