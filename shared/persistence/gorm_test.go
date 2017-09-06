package persistence

import (
	"testing"

	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/master/backend/config"
	repo_gorm "bitbucket.org/instinctools/gluten/shared/persistence/gorm"
	"github.com/jinzhu/gorm"
	assert "github.com/stretchr/testify/require"
)

func TestInitDb(t *testing.T) {
	conf := config.GlobalConfig.DB.Connection
	_, err := gorm.Open("postgres", conf.URL)
	assert.NoError(t, err)
}

func TestGormExecutionsRepo(t *testing.T) {
	exec_repo := repo_gorm.GetExecutionsRepo

	//1. test create object
	execution := core.Execution{
		ID:     "1",
		Status: "Status",
	}
	exec_repo.Create(&execution)

	//2. test getters
	db_execution := exec_repo.GetById("1")
	assert.NotEmpty(t, db_execution)

	//check pagination
	count_executions := len(exec_repo.Get(1, 0))
	assert.Equal(t, 1, count_executions)

	//3. test update object
	newStatusName := "NewStatus"
	db_execution.Status = newStatusName
	exec_repo.Update(db_execution)

	db_execution = exec_repo.GetById("1")
	assert.Equal(t, db_execution.Status, newStatusName)

	//4. test delete
	exec_repo.Delete(db_execution)
	db_execution = exec_repo.GetById("1")
	assert.Empty(t, db_execution)
}
