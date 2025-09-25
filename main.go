package main

import (
	"fmt"
	"task_tracker/storage"
	"task_tracker/tasks"
)

func main() {
	taskStorage, err := storage.NewFileStorage[tasks.Task]("tasks.json")
	task1 := tasks.NewTask("but cheese")
	if err != nil {
		fmt.Printf("error %v", err)
	}
	taskStorage.Save(task1)
}
