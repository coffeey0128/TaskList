package routes

import (
	"example.com/m/config"
	"example.com/m/driver"
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
