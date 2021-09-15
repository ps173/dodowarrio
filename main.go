package main

import (
	"flag"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	Name string
}

func main() {
	// Initializing the db
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		log.Panic("You got a err")
	}
	setup(db)

	// Flags management
	useListTodo := flag.Bool("ls", false, "List all todos")
	useDeleteAllTodo := flag.Bool("da", false, "Deletes all todos")
	useDeleteSingle := flag.Bool("del", false, "Deletes single todos with provided id of todo")
	useAddTodo := flag.Bool("add", false, "Adds a todo with defined string")
	flag.Parse()

	if *useListTodo && flag.Arg(0) == "" {
		listTodo(db)
	} else if *useDeleteAllTodo && flag.Arg(0) == "" {
		deleteAll(db)
	} else if *useAddTodo && flag.Arg(0) != "" {
		newTodo(db, &Todos{Name: flag.Arg(0)})
	} else if *useDeleteSingle && flag.Arg(0) != "" {
		deleteTodo(flag.Arg(0), db)
	} else {
		log.Panicln("Wrong value see --help for usage")
	}

}
