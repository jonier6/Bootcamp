//To add tasks, use the command = go run main.go + task name
//To get task, use the command = go run main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"todo"
)

const fileName = ".todo.json"

func main() {

	var list todo.List

	
	err := list.Get(fileName)
	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	
	args := os.Args[1:]

	
	if len(args) > 0 {

		task := strings.Join(args, " ")

		list.Add(task)

		err := list.Save(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		return
	}

	
	for _, item := range list {
		fmt.Println(item.Task)
	}
}