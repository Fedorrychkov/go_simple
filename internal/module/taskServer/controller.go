package taskServer

import (
	"fmt"
	"log"
	"net/http"

	"gitlab.com/learn_api/v2/internal/helpers"
	"gitlab.com/learn_api/v2/internal/task"
)

type TaskServer struct {
	Store *task.TaskStore
}

func (server *TaskServer) GetAllTasksHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get all tasks at %s\n", req.URL.Path)

	allTasks := server.Store.GetAllTasks()

	helpers.RenderJson(w, allTasks)
}

func (server *TaskServer) TaskHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/task/" {
		// Request is plain "/task/", without trailing ID.
		if req.Method == http.MethodPost {
		} else if req.Method == http.MethodGet {
			server.GetAllTasksHandler(w, req)
		} else if req.Method == http.MethodDelete {
		} else {
			http.Error(w, fmt.Sprintf("expect method GET, DELETE or POST at /task/, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}
	log.Println("Handler task")
}
