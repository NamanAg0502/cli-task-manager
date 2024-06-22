package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/namanag0502/cli-task-manager/pkg/model"
)

func SaveTasks(tasks []model.Task) {
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

func LoadTasks(tasks []model.Task, lastId int64) []model.Task {
	file, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println("Error opening tasks file: ", err)
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	if err != nil {
		fmt.Println("Error decoding tasks: ", err)
		return nil
	}

	return tasks
}
