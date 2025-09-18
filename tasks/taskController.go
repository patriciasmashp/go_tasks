package tasks

type TaskController struct {
	tasks map[int]*Task
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

}

func (c TaskController) CreateTask(description string) {
	task := NewTask(description)
	c.tasks[len(c.tasks)] = task
}