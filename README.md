# Purpose

Just a basic todo app for command line

# What Does it do

- store the todos in sqlite3 instance

Commands for the cli app

```bash
# sample commmands

$ todos ls # lists all todos
----------------------
 1. Create log file - [id]
----------------------
 2. Nextjs Vimgore frontend - [id]
-----------------------

$ todos del id # deletes the id
todo of id deleted

$. todos add something # creates a new todo
added a todo do todos ls to check it

```

# Checklist

- [x] Initialize project structure
- [x] Making basic commands - add
- [x] Making basic commands - ls
- [ ] Making basic commands - del ( hard delete them from database )
- [ ] Making basic commands - update the todo
- [x] Initialize a sql instance
- [x] Save the todos to database
- [ ] Make this a bit more good looking
- [ ] Need a config to edit the ui maybe
