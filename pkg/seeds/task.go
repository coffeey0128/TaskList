package seeds

import (
	"TaskList/internal/task"
	"TaskList/models"
	"fmt"

	"gorm.io/gorm"
)

func CreateTask(orm *gorm.DB, task models.Task) error {
	return orm.Create(&task).Error
}

var count = 0

func AllTask() []Seed {
	var tasks []Seed
	tasks = append(tasks, taskGenerator(5, task.StatusComplete)...)
	tasks = append(tasks, taskGenerator(5, task.StatusIncomplete)...)
	count = 0
	return tasks
}

func taskGenerator(needCount, status int) []Seed {
	var seeds = []Seed{}
	for i := 1; i <= needCount; i++ {
		count++
		seeds = append(seeds, Seed{
			Name: fmt.Sprintf("Create Tasks - name = task %d", count),
			Run: func(orm *gorm.DB) error {
				count++
				return CreateTask(orm, models.Task{
					Name:   fmt.Sprintf("task %d", count),
					Status: int64(status),
				})
			},
		})
	}
	return seeds
}
