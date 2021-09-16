package main

import (
	"fmt"

	"github.com/fatih/color"
	"gorm.io/gorm"
)

// HACK: Don't Print Directly
var c = color.New(color.FgMagenta)
var completed = color.New(color.BgGreen).Add(color.FgBlack)

func setup(db *gorm.DB) {
	db.AutoMigrate(&Todos{})
}

func newTodo(db *gorm.DB, todo *Todos) {
	db.Create(todo)
	c.Printf("New Todo is created with Id %d\n", todo.ID)
}

func listTodo(db *gorm.DB) {
	// Remeber : always pass a struct to be filled.
	var todos []Todos
	db.Find(&todos)
	for i, y := range todos {
		if y.Status {
			completed.Printf("%d. %s - %d [%t]", i+1, y.Name, y.ID, y.Status)
			fmt.Println()
		} else {
			color.Yellow("%d. %s - %d [%t]\n", i+1, y.Name, y.ID, y.Status)
		}
	}
}

func deleteTodo(key string, db *gorm.DB) {
	db.Unscoped().Delete(&Todos{}, key)
	c.Printf("Deleted todo with ID %s Succesfully \n", key)
}

func deleteAll(db *gorm.DB) {
	var todos []Todos
	db.Find(&todos)
	for _, y := range todos {
		db.Unscoped().Delete(&Todos{}, y.ID)
	}
	c.Println("Deleted All todos Succesfully")
}

func updateTodo(status bool, key string, db *gorm.DB) {
	var todo Todos
	db.First(&todo, key)
	todo.Status = status
	db.Save(&todo)
	color.Green("Updated todo with id %s with status %t succesfully", key, todo.Status)
}
