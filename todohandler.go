package main

import (
	"github.com/fatih/color"
	"gorm.io/gorm"
)

// TODO: Don't Print Directly
var c = color.New(color.FgMagenta)

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
		color.Green("%d. %s - %d \n", i+1, y.Name, y.ID)
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
