package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task
var lastId int64

func main() {
	loadTasks()
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
			addTask(scanner)
		case "2":
			listTasks()
		case "3":
			markTaskAsDone(scanner)
		case "4":
			deleteTask(scanner)
		case "5":
			saveTasks()
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task description: ")
	scanner.Scan()
	title := scanner.Text()
	lastId++
	task := Task{
		Id:    lastId,
		Title: title,
		Done:  false,
	}
	tasks = append(tasks, task)
	fmt.Println("\nTask added successfully!")
}

func listTasks() {
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

func markTaskAsDone(scanner *bufio.Scanner) {
	fmt.Print("Enter task ID: ")
	scanner.Scan()
	taskId, _ := strconv.Atoi(scanner.Text())

	for i, task := range tasks {
		if task.Id == int64(taskId) {
			tasks[i].Done = true
			fmt.Println("\nTask marked as done successfully!")
			return
		}
	}
	fmt.Println("\nTask not found!")
}

func deleteTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task ID to delete: ")
	scanner.Scan()
	taskId, _ := strconv.Atoi(scanner.Text())
	for i, task := range tasks {
		if task.Id == int64(taskId) {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("\nTask deleted successfully!")
			return
		}
	}
	fmt.Println("\nTask not found!")
}

func saveTasks() {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error creating tasks file: ", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding tasks: ", err)
		return
	}
	fmt.Println("\nTasks saved successfully!")
}

func loadTasks() {
	file, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println("Error opening tasks file: ", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	if err != nil {
		fmt.Println("Error decoding tasks: ", err)
		return
	}

	if len(tasks) > 0 {
		lastId = tasks[len(tasks)-1].Id
	}
}
