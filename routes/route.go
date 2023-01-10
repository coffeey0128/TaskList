package routes

import (
	"TaskList/config"
	"TaskList/docs"
	"TaskList/driver"
	"TaskList/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorResponse())
	// init swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//generate by code-gen
	TaskV1(r)

	return r
}

func init() {
	config.InitEnv()
	driver.InitGorm()
}
