package task

import (
	mockTask "TaskList/mock/task"
	"TaskList/models"
	"TaskList/models/apireq"
	"TaskList/pkg/er"
	"TaskList/pkg/query_condition"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Generate from template
func TestServiceFindAll(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockTaskRepo := mockTask.NewMockRepository(ctl)
	srvTask := NewTaskService(mockTaskRepo)

	mockTasks := []*models.Task{
		{ //自行補充
			Id: 1,
		},
	}

	tests := []struct {
		name           string
		page           int64
		perPage        int64
		queryCondition query_condition.QueryCondition
	}{
		{ // test 1
			name:           "test1",
			page:           1,
			perPage:        10,
			queryCondition: query_condition.QueryCondition{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTaskRepo.EXPECT().FindAll(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockTasks, nil)
			mockTaskRepo.EXPECT().Count(gomock.Any()).Return(1, nil)
			_, err := srvTask.FindAll(&apireq.ListTask{Page: test.page, PerPage: test.perPage}, test.queryCondition)
			assert.Nil(t, err)
		})
	}
}

// Generate from template
func TestServiceFindOne(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockTaskRepo := mockTask.NewMockRepository(ctl)
	srvTask := NewTaskService(mockTaskRepo)

	mockTasks := []*models.Task{
		{ //自行補充
			Id: 1,
		},
		{ //自行補充
			Id: 0,
		},
	}

	tests := []struct {
		name  string
		req   *apireq.GetTaskDetail
		found int // 0 = not found, 1 = found
		err   error
	}{
		{ // test 1
			name: "test1",
			req: &apireq.GetTaskDetail{
				Id: 1,
			},
			found: 1,
			err:   nil,
		},
		{ // test 2 not found
			name: "test2",
			req: &apireq.GetTaskDetail{
				Id: 2,
			},
			found: 0,
			err:   &er.AppError{StatusCode: 404, Code: "500000", Msg: "Task not found.", CauseErr: error(nil)},
		},
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTaskRepo.EXPECT().FindOne(gomock.Any()).Return(mockTasks[i], test.found, nil)
			_, err := srvTask.FindOne(test.req)
			assert.Equal(t, test.err, err)
		})
	}
}

// Generate from template
func TestServiceCreate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockTaskRepo := mockTask.NewMockRepository(ctl)
	srvTask := NewTaskService(mockTaskRepo)

	tests := []struct {
		name      string
		req       *apireq.CreateTask
		duplicate int
		err       error
	}{
		{ // test 1
			name:      "test1",
			req:       &apireq.CreateTask{},
			duplicate: 0,
			err:       nil,
		},
		{ // test 2 duplicate
			name:      "test2",
			req:       &apireq.CreateTask{},
			duplicate: 1,
			err:       &er.AppError{StatusCode: 409, Code: "400410", Msg: "create Task duplicate error.", CauseErr: error(nil)},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTaskRepo.EXPECT().FindOne(gomock.Any()).Return(&models.Task{}, test.duplicate, nil)
			if test.duplicate == 0 {
				mockTaskRepo.EXPECT().Insert(gomock.Any()).Return(nil)
			}
			err := srvTask.Create(test.req)
			assert.Equal(t, test.err, err)
		})
	}
}

// Generate from template
func TestServiceUpdate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockTaskRepo := mockTask.NewMockRepository(ctl)
	srvTask := NewTaskService(mockTaskRepo)

	mockTasks := []*models.Task{
		{ //自行補充
			Id: 1,
		},
		{ //自行補充
			Id: 0,
		},
	}

	tests := []struct {
		name    string
		req     *apireq.UpdateTask
		found   int // 0 = not found, 1 = found
		findErr error
		err     error
	}{
		{ // test 1
			name:    "test1",
			req:     &apireq.UpdateTask{},
			found:   0,
			findErr: nil,
			err:     nil,
		},
		{ // test 2 duplicate
			name:    "test2",
			req:     &apireq.UpdateTask{},
			found:   1,
			findErr: gorm.ErrRecordNotFound,
			err:     &er.AppError{StatusCode: 404, Code: "500000", Msg: "Task not found.", CauseErr: (nil)},
		},
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTaskRepo.EXPECT().FindOne(gomock.Any()).Return(mockTasks[i], test.found, test.findErr)
			if test.found == 0 {
				// 如果可重複FindOne要刪掉(Generate from template)
				mockTaskRepo.EXPECT().FindOne(gomock.Any()).Return(&models.Task{}, 0, nil)
				mockTaskRepo.EXPECT().Update(gomock.Any()).Return(nil)
			}
			err := srvTask.Update(test.req)
			assert.Equal(t, test.err, err)
		})
	}
}

// Generate from template
func TestServiceDelete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockTaskRepo := mockTask.NewMockRepository(ctl)
	srvTask := NewTaskService(mockTaskRepo)

	mockTasks := []*models.Task{
		{ //自行補充
			Id: 1,
		},
		{ //自行補充
			Id: 0,
		},
	}

	tests := []struct {
		name    string
		req     *apireq.DeleteTask
		found   int // 0 = not found, 1 = found
		findErr error
		err     error
	}{
		{ // test 1
			name: "test1",
			req: &apireq.DeleteTask{
				Id: 1,
			},
			found:   1,
			findErr: nil,
			err:     nil,
		},
		{ // test 2 not found
			name: "test2",
			req: &apireq.DeleteTask{
				Id: 2,
			},
			found:   0,
			findErr: gorm.ErrRecordNotFound,
			err:     &er.AppError{StatusCode: 404, Code: "500000", Msg: "Task not found.", CauseErr: error(nil)},
		},
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTaskRepo.EXPECT().FindOne(gomock.Any()).Return(mockTasks[i], test.found, test.findErr)
			if test.found != 0 {
				mockTaskRepo.EXPECT().Delete(gomock.Any()).Return(nil)
			}
			err := srvTask.Delete(test.req)
			assert.Equal(t, test.err, err)
		})
	}
}
