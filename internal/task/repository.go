package task

import (
	"TaskList/models"
	"TaskList/pkg/query_condition"
	"gorm.io/gorm"
)

type Repository interface {
	Insert(*models.Task) (err error)
	FindAll(page, perPage int64, queryCondition query_condition.QueryCondition) (record []*models.Task, err error)
	FindOne(*models.Task) (record *models.Task, rowaffected int, err error)
	Update(condition *models.Task) (err error)
	Delete(condition *models.Task) (err error)
	Count(queryCondition query_condition.QueryCondition) (int, error)
}

type TaskRepo struct {
	orm *gorm.DB
}

func NewTaskRepo(orm *gorm.DB) Repository {
	return &TaskRepo{orm: orm}
}

// Generate from template
func (r *TaskRepo) Insert(condition *models.Task) (err error) {
	result := r.orm.Create(condition)
	if err = result.Error; err != nil {
		return err
	}
	return nil
}

// Generate from template
func (r *TaskRepo) FindAll(page, perPage int64, queryCondition query_condition.QueryCondition) (record []*models.Task, err error) {
	offset := (page - 1) * perPage
	result := r.orm.Where(queryCondition.ToSQL()).Offset(int(offset)).Limit(int(perPage)).Find(&record)
	if err = result.Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (r *TaskRepo) FindOne(condition *models.Task) (record *models.Task, rowaffected int, err error) {
	record = new(models.Task)
	result := r.orm.First(record, condition)
	return record, int(result.RowsAffected), result.Error
}

// Generate from template
func (r *TaskRepo) Update(condition *models.Task) (err error) {
	result := r.orm.Select("*").Omit("created_at").Updates(condition)
	if err = result.Error; err != nil {
		return err
	}
	return nil
}

// Generate from template
func (r *TaskRepo) Delete(condition *models.Task) (err error) {
	result := r.orm.Delete(condition)
	if err = result.Error; err != nil {
		return err
	}
	return nil
}

func (r *TaskRepo) Count(queryCondition query_condition.QueryCondition) (int, error) {
	var count int64
	err := r.orm.Model(&models.Task{}).Where(queryCondition.ToSQL()).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
