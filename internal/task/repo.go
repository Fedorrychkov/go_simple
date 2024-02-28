package task

func InitTaskRepo() *TaskStore {
	store := &TaskStore{}
	store.tasks = make(map[int]Task)
	store.nextId = 0

	return store
}

func (store *TaskStore) CreateTask(text string, tags []string) int {
	task := Task{
		ID:   store.nextId,
		Text: text,
		Tags: tags,
	}

	task.Tags = make([]string, len(tags))
	copy(task.Tags, tags)

	store.tasks[store.nextId] = task
	store.nextId++
	return task.ID
}

func (store *TaskStore) GetAllTasks() []Task {
	allTasks := make([]Task, 0, len(store.tasks))

	for _, value := range store.tasks {
		allTasks = append(allTasks, value)
	}

	return allTasks
}
