package task

import (
	"TaskList/models"
	"TaskList/models/apireq"
	"TaskList/models/apires"
	"TaskList/pkg/er"
	"TaskList/pkg/query_condition"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Service interface {
	FindAll(req *apireq.ListTask, queryCondition query_condition.QueryCondition) (*apires.ListTask, error)
	FindOne(req *apireq.GetTaskDetail) (*apires.Task, error)
	Create(req *apireq.CreateTask) (err error)
	Update(req *apireq.UpdateTask) (err error)
	Delete(req *apireq.DeleteTask) (err error)
}

type TaskService struct {
	repo Repository
}

func NewTaskService(repo Repository) Service {
	return &TaskService{repo: repo}
}

// Generate from template
func (s *TaskService) FindAll(req *apireq.ListTask, queryCondition query_condition.QueryCondition) (*apires.ListTask, error) {
	result, err := s.repo.FindAll(req.Page, req.PerPage, queryCondition)
	if err != nil {
		return nil, er.NewAppErr(500, er.UnknownError, "ListAll Task error.", err)
	}
	totalCount, err := s.repo.Count(queryCondition)
	if err != nil {
		return nil, er.NewAppErr(500, er.UnknownError, "Count Task error.", err)
	}
	// need to transform models to apires
	results := make([]apires.Task, 0)
	if err := copier.Copy(&results, &result); err != nil {
		return nil, er.NewAppErr(500, er.UnknownError, "copy result to *apires.Task error.", err)
	}
	response := &apires.ListTask{
		Tasks:       results,
		CurrentPage: int(req.Page),
		PerPage:     int(req.PerPage),
		Total:       int(totalCount),
	}
	return response, nil
}

func (s *TaskService) FindOne(req *apireq.GetTaskDetail) (*apires.Task, error) {
	condition := &models.Task{Id: req.Id}
	record, rows, err := s.repo.FindOne(condition)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, er.NewAppErr(500, er.UnknownError, "get Task error.", err)
	}
	if rows == 0 {
		return nil, er.NewAppErr(404, er.UnknownError, "Task not found.", err)
	}
	res := &apires.Task{}
	if err := copier.Copy(res, record); err != nil {
		return nil, er.NewAppErr(500, er.UnknownError, "copy result to *apires.Task error.", err)
	}
	return res, nil
}

// Generate from template
func (s *TaskService) Create(req *apireq.CreateTask) (err error) {
	condition := &models.Task{}
	if err := copier.Copy(condition, req); err != nil {
		return er.NewAppErr(500, er.UnknownError, "copy *apireq.Task to *apires.Task error.", err)
	}
	// 檢查是否已經存在
	_, rows, err := s.repo.FindOne(condition)
	if err != nil && err != gorm.ErrRecordNotFound {
		return er.NewAppErr(500, er.UnknownError, "get Task error.", err)
	}
	if rows > 0 {
		return er.NewAppErr(409, er.DataDuplicateError, "create Task duplicate error.", err)
	}
	err = s.repo.Insert(condition)
	if err != nil {
		return er.NewAppErr(500, er.UnknownError, "create Task db error.", err)
	}
	return nil
}

// Generate from template
func (s *TaskService) Update(req *apireq.UpdateTask) (err error) {
	condition := &models.Task{Id: req.Id}
	// 檢查 record 是否存在
	_, rows, err := s.repo.FindOne(condition)
	if err != nil && err != gorm.ErrRecordNotFound {
		return er.NewAppErr(500, er.UnknownError, "get Task error.", err)
	}
	if err == gorm.ErrRecordNotFound {
		return er.NewAppErr(404, er.UnknownError, "Task not found.", nil)
	}
	// 檢查更新後的是否存在
	// 如果可重複這段要刪掉(Generate from template)
	checkExist := new(models.Task)
	if err := copier.Copy(checkExist, req); err != nil {
		return er.NewAppErr(500, er.UnknownError, "copy *apireq.Task to *models.Task error.", err)
	}
	_, rows, err = s.repo.FindOne(checkExist)
	if err != nil && err != gorm.ErrRecordNotFound {
		return er.NewAppErr(500, er.UnknownError, "get Task error.", err)
	}
	if rows > 0 {
		return er.NewAppErr(409, er.DataDuplicateError, "update Task duplicate error.", err)
	}
	err = s.repo.Update(checkExist)
	if err != nil {
		return er.NewAppErr(500, er.UnknownError, "update Task db error.", err)
	}
	return nil
}

// Generate from template
func (s *TaskService) Delete(req *apireq.DeleteTask) (err error) {
	condition := &models.Task{Id: req.Id}
	_, _, err = s.repo.FindOne(condition)
	if err != nil && err != gorm.ErrRecordNotFound {
		return er.NewAppErr(500, er.UnknownError, "get Task error.", err)
	}
	if err == gorm.ErrRecordNotFound {
		return er.NewAppErr(404, er.UnknownError, "Task not found.", nil)
	}
	err = s.repo.Delete(condition)
	if err != nil {
		return er.NewAppErr(500, er.UnknownError, "delete Task error.", err)
	}
	return nil
}
