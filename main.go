package main

import (
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

	// TODO : Better Cli and flags
	flag := os.Args[1]
	argument := os.Args[2:]
	if flag == "new" && argument != nil {
		argStr := strings.Join(argument, " ")
		newTodo(argStr, db)
	}

	if flag == "ls" {
		listTodo(db)
	}
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&Todos{})
}

func newTodo(s string, db *gorm.DB) {
	db.Create(&Todos{Name: s, Status: false})
}

func listTodo(db *gorm.DB) {
	//TODO: Figure Out how to list these
}
