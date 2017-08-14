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

func CreateExecution(execution Execution) Execution{
	tx := db.Begin()
	err := tx.Create(&execution).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	
	return execution
}

func GetExecution(id uint) Execution {
	var execution Execution
	db.Preload("Result.Metrics").First(&execution, id)
	return execution
}

// FIXME: Only one delete operation is needed
func DeleteExecution(execution Execution) {
	db.Delete(&execution.Result.Metrics)
	db.Delete(&execution.Result)
	db.Delete(&execution)
}
