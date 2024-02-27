package task

type Repository interface {
	CreateTask(text string, tags []string) int
	// GetTask(id int) (Task, error)
	// DeleteTask(id int) error
	GetAllTasks() []Task
}
