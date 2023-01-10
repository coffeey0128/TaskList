package apires

type Task struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type ListTask struct {
	Result []Task `json:"result"`
}
