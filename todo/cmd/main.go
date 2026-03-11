package main

import (
	"fmt"
	"os"
	"flag"
	"todo"
)

const fileName = ".todo.json"

func main() {

	listFlag := flag.Bool("list", false, "list tasks")
	taskFlag := flag.String("task", "", "add task")
	completeFlag := flag.Int("complete", -1, "complete task")
	deleteFlag := flag.Int("delete", -1, "delete task")

	flag.Parse()

	var list todo.List

	
	err := list.Get(fileName)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}

	
	switch {

	case *listFlag:

		for _, task := range list {
			if !task.Done {
				fmt.Printf(
					"Title: %s, Done: %t, CreatedAt: %s, CompletedAt: %s\n",
					task.Task,
					task.Done,
					task.CreatedAt,
					task.CompletedAt,
				)
			}
		}

	case *completeFlag != -1:

		err := list.Complete(*completeFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		list.Save(fileName)

	case *deleteFlag != -1:

		err := list.Delete(*deleteFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		list.Save(fileName)

	case *taskFlag != "":

		list.Add(*taskFlag)

		err := list.Save(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	default:

		fmt.Println("error: no command provided")
		os.Exit(1)
	}
}