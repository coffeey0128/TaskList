package task

import (
	"TaskList/config"
	"TaskList/driver"
	"TaskList/models"
	"TaskList/pkg/query_condition"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// Generate from template
func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

// Generate from template
func setUp() {
	remoteBranch := os.Getenv("REMOTE_BRANCH")
	if remoteBranch == "" {
		// load env
		err := godotenv.Load(config.GetBasePath() + "/.env")
		if err != nil {
			log.Panicln(err)
		}
	}
}

// Generate from template
func TestRepoInsert(t *testing.T) {
	gormEngine := driver.InitGorm()
	repo := NewTaskRepo(gormEngine)
	//自行填入struct
	condition := &models.Task{}
	// 自行接response測試
	err := repo.Insert(condition)
	assert.Nil(t, err)

	// tear down
	_ = gormEngine.Delete(models.Task{}, condition.Id)
}

// Generate from template
func TestRepoFindAll(t *testing.T) {
	gormEngine := driver.InitGorm()
	repo := NewTaskRepo(gormEngine)
	var queryCondition query_condition.QueryCondition
	// 自行接response測試
	_, err := repo.FindAll(1, 10, queryCondition)
	assert.Nil(t, err)
}

// Generate from template
func TestRepoFindOne(t *testing.T) {
	gormEngine := driver.InitGorm()
	repo := NewTaskRepo(gormEngine)
	// 自行接response測試
	_, _, err := repo.FindOne(&models.Task{Id: 1})
	assert.Nil(t, err)
}

// Generate from template
func TestRepoUpdate(t *testing.T) {
	gormEngine := driver.InitGorm()
	repo := NewTaskRepo(gormEngine)
	//自行填入struct
	condition := &models.Task{}
	// 自行接response測試
	err := repo.Update(condition)
	assert.Nil(t, err)
}

// Generate from template
func TestRepoDelete(t *testing.T) {
	gormEngine := driver.InitGorm()
	repo := NewTaskRepo(gormEngine)
	//自行填入struct
	condition := &models.Task{}
	_ = gormEngine.Create(condition)
	// 自行接response測試
	err := repo.Delete(condition)
	assert.Nil(t, err)
}

// Generate from template
func TestRepoCount(t *testing.T) {
	gormEngine := driver.InitGorm()
	repo := NewTaskRepo(gormEngine)
	var queryCondition query_condition.QueryCondition
	// 自行接response測試
	_, err := repo.Count(queryCondition)
	assert.Nil(t, err)
}
