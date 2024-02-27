package main

import (
	"fmt"

	"gitlab.com/learn_api/v2/internal/task"
)

func main() {
	store := task.InitTaskRepo()

	store.CreateTask("test", []string{"tag1", "tag2"})
	store.CreateTask("test", []string{"tag4", "tag3"})

	fmt.Println("Test api", store.GetAllTasks())
}
