package routes

import (
	"TaskList/config"
	"TaskList/driver"
	"TaskList/middleware"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorResponse())
	return r
}

func init() {
	config.InitEnv()
	driver.InitGorm()
}
