package repositories

import (
	"bitbucket.org/instinctools/gluten/shared/persistence"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	DB = persistence.InitDb()
}

func CreateExecution(execution persistence.Execution) persistence.Execution {
	tx := DB.Begin()
	err := tx.Create(&execution).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	return execution
}

func GetExecutions() []persistence.Execution {
	var executions []persistence.Execution
	DB.Preload("Result.Metrics").Find(&executions)
	return executions

}

func GetExecutionById(id uint) persistence.Execution {
	var execution persistence.Execution
	DB.Preload("Result.Metrics").First(&execution, id)
	return execution
}

func DeleteExecution(execution persistence.Execution) {
	DB.Delete(&execution)
}

func CloseConnection() {
	DB.Close()
}
