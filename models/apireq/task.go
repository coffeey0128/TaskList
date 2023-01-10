package apireq

type Task struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type ListTask struct {
	Page    int `form:"page" binding:"required,min=1" json:"page"`
	PerPage int `form:"per_page" binding:"required" json:"per_page"`
}

type ListTaskQueryCondition struct {
	NameEq   string `json:"name" form:"name_eq"`
	StatusEq int    `json:"status" form:"status_eq"`
}

type GetTaskDetail struct {
	Id int `uri:"id" form:"id" binding:"required" json:"id"`
}

type CreateTask struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type UpdateTask struct {
	Id     int    `uri:"id" form:"id" binding:"required" json:"id"`
	Name   string `json:"name" binding:"required"`
	Status int    `json:"status" binding:"required"`
}

type DeleteTask struct {
	Id int `uri:"id" form:"id" binding:"required" json:"id"`
}
