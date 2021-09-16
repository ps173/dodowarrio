package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fatih/color"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	Name   string
	Status bool
}

func main() {
	// Initializing the db in home dir
	dirname, _ := os.UserHomeDir()
	p := filepath.Join(dirname, "main.db")

	db, err := gorm.Open(sqlite.Open(p), &gorm.Config{})
	if err != nil {
		log.Panic("You got a err")
	}
	setup(db)

	// Flags management
	useListTodo := flag.Bool("ls", false, "List all todos")
	useDeleteAllTodo := flag.Bool("da", false, "Deletes all todos")
	useDeleteSingle := flag.Bool("del", false, "Deletes single todos with provided id of todo")
	useAddTodo := flag.Bool("add", false, "Adds a todo with defined string")
	useUpdateTodo := flag.Bool("ud", false, "Updates a todo with given id and status [ remember status should be boolean ]")
	flag.Parse()

	if *useListTodo && flag.Arg(0) == "" {
		listTodo(db)
	} else if *useDeleteAllTodo && flag.Arg(0) == "" {
		deleteAll(db)
	} else if *useAddTodo && flag.Arg(0) != "" {
		newTodo(db, &Todos{Name: flag.Arg(0), Status: false})
	} else if *useDeleteSingle && flag.Arg(0) != "" {
		deleteTodo(flag.Arg(0), db)
	} else if *useUpdateTodo && flag.Arg(1) != "" {
		status, _ := strconv.ParseBool(flag.Arg(1))
		updateTodo(status, flag.Arg(0), db)
	} else {
		color.Red("Wrong value see --help for usage")
	}

}
