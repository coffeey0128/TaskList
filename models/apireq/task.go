package apireq

type Task struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Status int64  `json:"status"`
}

type ListTask struct {
	Page    int64 `form:"page" binding:"required,min=1" json:"page"`
	PerPage int64 `form:"per_page" binding:"required" json:"per_page"`
}

type ListTaskQueryCondition struct {
	NameEq   string `json:"name" form:"name_eq"`
	StatusEq int64  `json:"status" form:"status_eq"`
}

type GetTaskDetail struct {
	Id int64 `uri:"id" form:"id" binding:"required" json:"id"`
}

type CreateTask struct {
	Name   string `json:"name"`
	Status int64  `json:"status"`
}

type UpdateTask struct {
	Id     int64  `uri:"id" form:"id" binding:"required" json:"id"`
	Name   string `json:"name" binding:"required"`
	Status int64  `json:"status" binding:"required"`
}

type DeleteTask struct {
	Id int64 `uri:"id" form:"id" binding:"required" json:"id"`
}
