package tasks

import (
	"fmt"
	"task_tracker/storage"
)

type TaskController struct {
	tasks    map[int]*Task
	latestId int
	storage  storage.Storage[Task]
}

func NewTaskController(tasks []Task) *TaskController {
	controller := &TaskController{
		tasks: make(map[int]*Task),
	}
	for _, task := range tasks {
		controller.tasks[task.ID] = &task
	}
	return controller
}

func (c TaskController) save() {
	taskStorage, err := storage.NewFileStorage[Task]("test.json")
	if err != nil {
		fmt.Printf("error on save Task%v", err)
		return
	}

	for _, task := range c.tasks {
		taskStorage.Save(task)
	}
}

func (c TaskController) CreateTask(description string) {
	task := NewTask(description)
	c.tasks[len(c.tasks)] = task
}
