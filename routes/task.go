package routes

import (
	"TaskList/api"

	"github.com/gin-gonic/gin"
)

func TaskV1(r *gin.Engine) {
	v1 := r.Group("")

	// from generator
	v1.GET("/tasks", api.ListTask)
	v1.GET("/tasks/:id", api.GetTaskDetail)
	v1.POST("/tasks", api.CreateTask)
	v1.PUT("/tasks/:id", api.UpdateTask)
	v1.DELETE("/tasks/:id", api.DeleteTask)
}
