package apires

type Task struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Status int64  `json:"status"`
}

type ListTask struct {
	Tasks       []Task `json:"tasks"`
	CurrentPage int    `json:"current_page"`
	PerPage     int    `json:"per_page"`
	Total       int    `json:"total"`
}
