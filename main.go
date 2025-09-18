package main

import (
	"fmt"
	"task_tracker/cli"
	"task_tracker/tasks"
)

func main() {
	cli := cli.ParseArgs()
	cli.PrintArgs()
	task := tasks.NewTask("test")
	fmt.Println(task)
}