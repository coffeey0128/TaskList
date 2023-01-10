//go:build wireinject
// +build wireinject

package api

import (
	"TaskList/driver"
	task "TaskList/internal/task"

	"github.com/google/wire"
)

var gormSet = wire.NewSet(driver.InitGorm)
var TaskServiceSet = wire.NewSet(task.NewTaskService, task.NewTaskRepo)

func BuildTaskSrv() task.Service {
	wire.Build(TaskServiceSet, gormSet)
	return &task.TaskService{}
}
