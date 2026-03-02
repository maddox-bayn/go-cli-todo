package main

import (
	"fmt"
	"go-cli-todo/todo"
	"os"
	//"github.com/maddox-bayn/go-cli-todo/todo"
)

func main() {
	storage := todo.NewFileStorage("todos.json")
	service := todo.NewTodoService(storage)

	if len(os.Args) < 2 {
		fmt.Println("usage todo [add][list]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: add \"task title\"")
			return
		}
		err := service.AddTodo(os.Args[2])
		if err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println("Todo added successfully!")
	default:
		fmt.Println("unknown command")
	}
}
