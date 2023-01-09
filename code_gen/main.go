package main

import (
	"TaskList/code_gen/generator"
	"TaskList/config"
	"TaskList/driver"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

var (
	projectName  string
	tableName    string
	internalPath string
	need         string
	genRepo      bool
	genSrv       bool
	genModel     bool
	genRoute     bool
	genApi       bool
)

func main() {
	for tableName == "" {
		fmt.Print("Please enter tableName: ")
		fmt.Println("mysql 已經新增的表名, 會根據 mysql table schema 產生對應內容")
		fmt.Scanln(&tableName)
	}
	fmt.Println()

	for projectName == "" {
		fmt.Print("Please enter projectName: ")
		fmt.Println("根據輸入的專案名產生 import 路徑，例如: TaskList")
		fmt.Scanln(&projectName)
	}
	fmt.Println()

	fmt.Print("Please enter path: ")
	fmt.Println("在internal底下的資料夾, 前不用加 / ，最後要補/ ex : post/  —> internal/post/")
	fmt.Scanln(&internalPath)
	fmt.Println()

	fmt.Print("Need repository[y/n]? ")
	fmt.Println("會在專案目錄 internal 產生對應的 table_name repository.go")
	fmt.Scanln(&need)
	if strings.ToLower(need) == "y" {
		genRepo = true
	}
	need = ""
	fmt.Println()

	fmt.Print("Need service[y/n]? ")
	fmt.Println("會在專案目錄 internal 產生對應的 table_name service.go")
	fmt.Scanln(&need)
	if strings.ToLower(need) == "y" {
		genSrv = true
	}
	need = ""
	fmt.Println()

	fmt.Print("Need model[y/n]? ")
	fmt.Println("會在專案目錄 models 產生對應的 table_name {table_name}.go")
	fmt.Scanln(&need)
	if strings.ToLower(need) == "y" {
		genModel = true
	}
	need = ""
	fmt.Println()

	fmt.Print("Need route[y/n]? ")
	fmt.Println("會在專案目錄 routes 產生對應的 {table_name}.go")
	fmt.Scanln(&need)
	if strings.ToLower(need) == "y" {
		genRoute = true
	}
	need = ""
	fmt.Println()

	fmt.Print("Need api[y/n]? ")
	fmt.Println("會在專案目錄 api 下產生對應的 {table_name}_api.go")
	fmt.Scanln(&need)
	if strings.ToLower(need) == "y" {
		genApi = true
	}
	need = ""

	godotenv.Load(config.GetBasePath() + "/.env")
	gorm := driver.InitGorm()
	generator := &generator.Generator{
		TableName:        tableName,
		InternalPath:     internalPath,
		ProjectName:      projectName,
		ModelPackageName: "models",
		GenRepo:          genRepo,
		GenService:       genSrv,
		GenModel:         genModel,
		GenApi:           genApi,
		GenRoute:         genRoute,
		Db:               gorm,
	}
	generator.GetColumns()
	generator.LoadTemplates()
	err := generator.Execute()
	if err != nil {
		panic(err)
	}
}
