package task

import (
	"sync"
)

type TaskStore struct {
	sync.Mutex

	tasks  map[int]Task
	nextId int
}
