package persistence

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = InitDb()
}

func CloseConnection() {
	db.Close()
}

func CreateExecution(execution Execution) Execution {
	tx := db.Begin()
	err := tx.Create(&execution).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	return execution
}

func GetExecutions() []Execution {
	var executions []Execution
	db.Preload("Result.Metrics").Find(&executions)
	return executions

}

func GetResults(id uint) []ExecutionResult {
	var results []ExecutionResult
	GetExecution(id)
	db.Preload("Metrics").Find(&results)
	return results

}

func GetExecution(id uint) Execution {
	var execution Execution
	db.Preload("Result.Metrics").First(&execution, id)
	return execution
}

func DeleteExecution(execution Execution) {
	db.Delete(&execution)
}
