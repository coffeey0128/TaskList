package routes

import (
	"TaskList/config"
	"TaskList/driver"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	return r
}

func init() {
	config.InitEnv()
	driver.InitGorm()
}
