package task

type Task struct {
	ID   int      `json:"id"`
	Text string   `json:"text"`
	Tags []string `json:"tags"`
}
