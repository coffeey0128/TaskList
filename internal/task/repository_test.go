package task

import (
	"TaskList/config"
	"TaskList/driver"
	"TaskList/models"
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
	task := &models.Task{
		Name: "買早餐",
	}
	err := repo.Insert(task)

	assert.Nil(t, err)
	assert.NotEqual(t, task.Id, 0)

	// tear down
	_ = gormEngine.Delete(models.Task{}, task.Id)
}

// Generate from template
func TestRepoFindAll(t *testing.T) {
	gormEngine := driver.InitGorm()
	repo := NewTaskRepo(gormEngine)
	tasks, err := repo.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 10, len(tasks))
}

// Generate from template
func TestRepoFindOne(t *testing.T) {
	gormEngine := driver.InitGorm()
	repo := NewTaskRepo(gormEngine)
	_, _, err := repo.FindOne(&models.Task{Id: 1})
	assert.Nil(t, err)

	_, _, err = repo.FindOne(&models.Task{Id: 1000000})
	assert.NotNil(t, err)
}

// Generate from template
func TestRepoUpdate(t *testing.T) {
	gormEngine := driver.InitGorm()
	repo := NewTaskRepo(gormEngine)
	updatedTask := &models.Task{
		Id:     1,
		Name:   "Updated task",
		Status: StatusComplete,
	}
	err := repo.Update(updatedTask)
	assert.Nil(t, err)

	expectedTask := new(models.Task)
	result := gormEngine.First(expectedTask, &models.Task{Id: updatedTask.Id})
	assert.Nil(t, result.Error)
	assert.NotZero(t, result.RowsAffected)
	assert.Equal(t, expectedTask.Name, updatedTask.Name)
	assert.Equal(t, expectedTask.Status, updatedTask.Status)

	// teardown
	result = gormEngine.Select("*").Omit("created_at").Updates(&models.Task{
		Id:     1,
		Name:   "task 1",
		Status: StatusIncomplete,
	})
	assert.Nil(t, result.Error)
	assert.NotZero(t, result.RowsAffected)
}

// Generate from template
func TestRepoDelete(t *testing.T) {
	gormEngine := driver.InitGorm()
	repo := NewTaskRepo(gormEngine)
	//自行填入struct
	condition := &models.Task{
		Name:   "Test delete task",
		Status: 0,
	}
	result := gormEngine.Create(condition)
	assert.Nil(t, result.Error)

	err := repo.Delete(condition)
	assert.Nil(t, err)

	result = gormEngine.First(condition)
	assert.NotNil(t, result.Error)
	assert.Zero(t, result.RowsAffected)
}
