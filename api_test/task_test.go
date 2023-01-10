package api_test

import (
	"TaskList/internal/task"
	"TaskList/models/apireq"
	"TaskList/models/apires"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAllTask(t *testing.T) {
	t.Run("List All Task", func(t *testing.T) {
		path := fmt.Sprintf("/tasks?name_like=%s&page=%d&per_page=%d", "", 1, 10)
		request, _ := http.NewRequest(http.MethodGet, path, nil)
		response := httptest.NewRecorder()
		route.ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)

		var result apires.ListTask
		bodyBytes := response.Body.Bytes()
		fmt.Printf("%s", bodyBytes)
		err := json.Unmarshal(bodyBytes, &result)
		if err != nil {
			panic(err)
		}
		assert.Len(t, result.Result, 10)
	})
}

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name         string
		body         *apireq.CreateTask
		responseCode int
	}{
		{
			name: "create test StatusIncomplete",
			body: &apireq.CreateTask{
				Name:   "create task 1",
				Status: task.StatusIncomplete,
			},
			responseCode: 201,
		},
		{
			name: "create test StatusComplete",
			body: &apireq.CreateTask{
				Name:   "create task 2",
				Status: task.StatusComplete,
			},
			responseCode: 201,
		},
		{
			name: "create test StatusUnknown",
			body: &apireq.CreateTask{
				Name:   "create task 3",
				Status: 2,
			},
			responseCode: 400,
		},
		{
			name: "create test without name",
			body: &apireq.CreateTask{
				Name:   "",
				Status: 1,
			},
			responseCode: 400,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Create Task %s", test.name), func(t *testing.T) {
			body, err := json.Marshal(test.body)
			if err != nil {
				t.Fatal(err)
			}
			request, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
			response := httptest.NewRecorder()
			request.Header.Add("Content-type", "application/json")
			route.ServeHTTP(response, request)

			assert.Equal(t, test.responseCode, response.Code)
		})
	}
}

func TestUpdateTask(t *testing.T) {
	statusComplete := task.StatusComplete
	statusIncomplete := task.StatusIncomplete
	statusUnknown := 3
	tests := []struct {
		name         string
		id           int
		body         *apireq.UpdateTask
		response     string
		responseCode int
	}{
		{
			name: "test 1",
			id:   1,
			body: &apireq.UpdateTask{
				Name:   "update test",
				Status: &statusComplete,
			},
			response:     "{\"id\":1,\"name\":\"update test\",\"status\":1}",
			responseCode: 200,
		},
		{
			name: "test 2",
			id:   1,
			body: &apireq.UpdateTask{
				Name:   "update test 2",
				Status: &statusIncomplete,
			},
			response:     "{\"id\":1,\"name\":\"update test 2\",\"status\":0}",
			responseCode: 200,
		},
		{
			name: "test 3",
			id:   1,
			body: &apireq.UpdateTask{
				Name:   "update task 3",
				Status: &statusUnknown,
			},
			response:     "{\"code\":\"400400\",\"message\":\"Key: 'UpdateTask.Status' Error:Field validation for 'Status' failed on the 'oneof' tag\"}",
			responseCode: 400,
		},
		{
			name: "test 4",
			id:   1,
			body: &apireq.UpdateTask{
				Name: "update task 4",
			},
			response:     "{\"code\":\"400400\",\"message\":\"Key: 'UpdateTask.Status' Error:Field validation for 'Status' failed on the 'required' tag\"}",
			responseCode: 400,
		},
		{
			name: "test 5",
			id:   1,
			body: &apireq.UpdateTask{
				Status: &statusComplete,
			},
			response:     "{\"code\":\"400400\",\"message\":\"Key: 'UpdateTask.Name' Error:Field validation for 'Name' failed on the 'required' tag\"}",
			responseCode: 400,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Update Task %s", test.name), func(t *testing.T) {
			body, err := json.Marshal(test.body)
			if err != nil {
				t.Fatal(err)
			}
			path := fmt.Sprintf("/tasks/%d", 1)
			request, _ := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(body))
			response := httptest.NewRecorder()
			request.Header.Add("Content-type", "application/json")
			route.ServeHTTP(response, request)

			assert.Equal(t, test.responseCode, response.Code)
			assert.Equal(t, test.response, response.Body.String())
		})
	}
}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		id           int
		name         string
		response     string
		responseCode int
	}{
		{
			id:           1,
			name:         "test 1",
			response:     "{}",
			responseCode: 200,
		},
		{
			id:           1000,
			name:         "test 2",
			response:     "{\"code\":\"400404\",\"message\":\"Task not found.\"}",
			responseCode: 400,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Delete Task %s", test.name), func(t *testing.T) {
			path := fmt.Sprintf("/tasks/%d", test.id)
			request, _ := http.NewRequest(http.MethodDelete, path, nil)
			response := httptest.NewRecorder()
			request.Header.Add("Content-type", "application/json")
			route.ServeHTTP(response, request)
			assert.Equal(t, test.responseCode, response.Code)
			assert.Equal(t, test.response, response.Body.String())
		})
	}
}
