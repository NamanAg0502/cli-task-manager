package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/namanag0502/cli-task-manager/pkg/handler"
	"github.com/namanag0502/cli-task-manager/pkg/model"
	"github.com/namanag0502/cli-task-manager/pkg/utils"
)

var tasks []model.Task
var lastId int64

func main() {
	tasks = utils.LoadTasks(tasks, lastId)
	if len(tasks) > 0 {
		lastId = tasks[len(tasks)-1].Id
	} else {
		lastId = 0
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("===== WELCOME TO TASKLIST =====")
	for {
		fmt.Println("================================")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Println("================================")
		print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()
		fmt.Println("================================")

		switch choice {
		case "1":
			tasks = handler.AddTask(scanner, lastId, tasks)
		case "2":
			handler.ListTasks(tasks)
		case "3":
			b, i := handler.MarkTaskAsDone(scanner, tasks)
			tasks[i].Done = b
		case "4":
			tasks = handler.DeleteTask(scanner, tasks)
		case "5":
			utils.SaveTasks(tasks)
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}
}
