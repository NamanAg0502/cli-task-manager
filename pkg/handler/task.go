package handler

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/namanag0502/cli-task-manager/pkg/model"
)

func AddTask(scanner *bufio.Scanner, lastId int64, tasks []model.Task) []model.Task {
	fmt.Print("Enter task description: ")
	scanner.Scan()
	title := scanner.Text()
	lastId++
	task := model.Task{
		Id:    lastId,
		Title: title,
		Done:  false,
	}
	fmt.Println("\nTask added successfully!")
	return append(tasks, task)
}

func ListTasks(tasks []model.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found!")
	} else {
		fmt.Println("ID\tDone\tTitle")
	}

	for _, task := range tasks {
		fmt.Printf("%d\t%t\t%s\n", task.Id, task.Done, task.Title)
	}
	fmt.Println("\nTask list completed successfully!")
}

func MarkTaskAsDone(scanner *bufio.Scanner, tasks []model.Task) (bool, int) {
	fmt.Print("Enter task ID: ")
	scanner.Scan()
	taskId, _ := strconv.Atoi(scanner.Text())

	for i, task := range tasks {
		if task.Id == int64(taskId) {
			tasks[i].Done = true
			fmt.Println("\nTask marked as done successfully!")
			return true, i
		}
	}
	fmt.Println("\nTask not found!")
	return false, 0
}

func DeleteTask(scanner *bufio.Scanner, tasks []model.Task) []model.Task {
	fmt.Print("Enter task ID to delete: ")
	scanner.Scan()
	taskId, _ := strconv.Atoi(scanner.Text())
	for i, task := range tasks {
		if task.Id == int64(taskId) {
			fmt.Println("\nTask deleted successfully!")
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	fmt.Println("\nTask not found!")
	return nil
}
