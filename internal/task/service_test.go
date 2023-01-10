package task

import (
	mockTask "TaskList/mock/task"
	"TaskList/models"
	"TaskList/models/apireq"
	"TaskList/pkg/er"
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
		{
			Id:     1,
			Name:   "mock test 1",
			Status: StatusIncomplete,
		},
		{
			Id:     2,
			Name:   "mock test 1",
			Status: StatusComplete,
		},
	}

	t.Run("Test Find All Tasks", func(t *testing.T) {
		mockTaskRepo.EXPECT().FindAll().Return(mockTasks, nil)
		result, err := srvTask.FindAll()
		assert.Nil(t, err)
		assert.Equal(t, len(mockTasks), len(result.Result))
		for i := range result.Result {
			assert.Equal(t, mockTasks[i].Name, result.Result[i].Name)
			assert.Equal(t, mockTasks[i].Status, result.Result[i].Status)
		}
	})
}

// Generate from template
func TestServiceCreate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockTaskRepo := mockTask.NewMockRepository(ctl)
	srvTask := NewTaskService(mockTaskRepo)

	tests := []struct {
		name string
		req  *apireq.CreateTask
		err  error
	}{
		{ // test 1
			name: "test1",
			req: &apireq.CreateTask{
				Name:   "test 1",
				Status: StatusComplete,
			},
			err: nil,
		},
		{ // test 2 duplicate
			name: "test2",
			req: &apireq.CreateTask{
				Name:   "test2",
				Status: StatusIncomplete,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTaskRepo.EXPECT().Insert(gomock.Any()).Return(nil)
			res, err := srvTask.Create(test.req)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.req.Name, res.Name)
			assert.Equal(t, test.req.Status, res.Status)
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
		{
			Id:     1,
			Name:   "test task 1",
			Status: StatusComplete,
		},
		{ // for not found task
			Id: 2,
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
			name: "test1",
			req: &apireq.UpdateTask{
				Id:     1,
				Name:   "update test 1",
				Status: StatusIncomplete,
			},
			found:   1,
			findErr: nil,
			err:     nil,
		},
		{ // test 2 not found
			name: "test2",
			req: &apireq.UpdateTask{
				Id:     1000000,
				Name:   "update test 2",
				Status: StatusComplete,
			},
			found:   0,
			findErr: gorm.ErrRecordNotFound,
			err:     &er.AppError{StatusCode: 404, Code: "500000", Msg: "Task not found.", CauseErr: (nil)},
		},
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTaskRepo.EXPECT().FindOne(gomock.Any()).Return(mockTasks[i], test.found, test.findErr)
			if test.found == 1 {
				mockTaskRepo.EXPECT().Update(gomock.Any()).Return(nil)
			}
			res, err := srvTask.Update(test.req)
			assert.Equal(t, test.err, err)
			if err == nil {
				assert.Equal(t, test.req.Name, res.Name)
				assert.Equal(t, test.req.Status, res.Status)
			}
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
		{
			Id:     1,
			Name:   "test 1",
			Status: StatusComplete,
		},
		{
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
				Id: 10000,
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
