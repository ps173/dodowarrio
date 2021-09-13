package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Todos struct {
	gorm.Model
	Name   string
	Status bool
}

func main() {
	// Logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panic("You got a err")
	}
	setup(db)

	flag := os.Args[1]
	argument := os.Args[2:]
	if flag == "new" && argument != nil {
		argStr := strings.Join(argument, " ")
		todo := Todos{Name: argStr, Status: false}
		newTodo(db, &todo)
	}

	if flag == "ls" {
		listTodo(db)
	}
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&Todos{})
}

func newTodo(db *gorm.DB, todo *Todos) {
	db.Create(todo)
}

func listTodo(db *gorm.DB) {
	// Remeber : always pass a struct to be filled.
	var todos Todos
	db.Find(&todos)
	fmt.Println(todos)
	// db.First(&todos)
	// fmt.Printf("Todo:%s\nCreated-At:%s\n", todos.Name, todos.CreatedAt)
}
