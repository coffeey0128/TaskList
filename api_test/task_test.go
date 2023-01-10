package api_test

import (
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
	})
}

func TestGetTaskDetail(t *testing.T) {
	t.Run("Get Task detail", func(t *testing.T) {
		path := fmt.Sprintf("/tasks/%d", 1)
		request, _ := http.NewRequest(http.MethodGet, path, nil)
		response := httptest.NewRecorder()
		route.ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)

		var result apires.Task
		bodyBytes := response.Body.Bytes()
		fmt.Printf("%s", bodyBytes)
		err := json.Unmarshal(bodyBytes, &result)
		if err != nil {
			panic(err)
		}
	})
}

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name         string
		body         *apireq.CreateTask
		response     string
		responseCode int
	}{
		{
			name:         "test 1",
			body:         &apireq.CreateTask{}, // 必須填值
			response:     "{}",
			responseCode: 200,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Create Task %s", test.name), func(t *testing.T) {
			body, err := json.Marshal(test.body)
			if err != nil {
				t.Fatal(err)
			}
			path := fmt.Sprintf("/tasks")
			request, _ := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(body))
			response := httptest.NewRecorder()
			request.Header.Add("Content-type", "application/json")
			route.ServeHTTP(response, request)

			assert.Equal(t, test.responseCode, response.Code)
			assert.Equal(t, test.response, response.Body.String())
		})
	}
}

func TestUpdateTask(t *testing.T) {
	tests := []struct {
		name         string
		body         *apireq.UpdateTask
		response     string
		responseCode int
	}{
		{
			name:         "test 1",
			body:         &apireq.UpdateTask{}, // 必須填值
			response:     "{}",
			responseCode: 200,
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
		name         string
		response     string
		responseCode int
	}{
		{
			name:         "test 1",
			response:     "{}",
			responseCode: 200,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Delete Task %s", test.name), func(t *testing.T) {
			path := fmt.Sprintf("/tasks/%d", 1)
			request, _ := http.NewRequest(http.MethodDelete, path, nil)
			response := httptest.NewRecorder()
			request.Header.Add("Content-type", "application/json")
			route.ServeHTTP(response, request)
			assert.Equal(t, test.responseCode, response.Code)
			assert.Equal(t, test.response, response.Body.String())
		})
	}
}
