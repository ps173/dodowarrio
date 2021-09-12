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

type Todo struct {
	gorm.Model
	Name   string
	Status bool
}

func main() {
	// Logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
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
		newTodo(argStr, db)
	}
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&Todo{})
}

func newTodo(s string, db *gorm.DB) {
	db.Create(&Todo{Name: s, Status: false})
}

func listTodo(db *gorm.DB) {

}
