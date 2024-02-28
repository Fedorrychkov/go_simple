package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/learn_api/v2/internal/module/taskServer"
	"gitlab.com/learn_api/v2/internal/task"
)

func NewTaskServer() *taskServer.TaskServer {
	store := task.InitTaskRepo()

	store.CreateTask("test", []string{"tag1", "tag2"})

	return &taskServer.TaskServer{Store: store}
}

func main() {
	mux := http.NewServeMux()

	server := NewTaskServer()
	mux.HandleFunc("/task/", server.TaskHandler)

	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), mux))

}
