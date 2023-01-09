package api_test

import (
	"TaskList/routes"
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var route *gin.Engine

func setup() {
	route = routes.Init()
	gin.SetMode(gin.TestMode)
	fmt.Println("Before all tests")
}
func teardown() {
	fmt.Println("After all tests")
}
func TestMain(m *testing.M) {
	setup()
	fmt.Println("Test begins....")
	code := m.Run()
	teardown()
	os.Exit(code)
}
