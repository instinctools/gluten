package gorm

import (
	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=gluten sslmode=disable password=1")
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("gorm.Open error")
	}
	db.LogMode(true)
	return db
}

//ExecutionRepoAdapter
/*func CreateExecution(execution persistence.Execution) persistence.Execution {
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

func UpdateExecution(execution persistence.Execution) {
	DB.Update(execution)
}

//ResultRepoAdapter
func GetResults(id uint) []persistence.ExecutionResult {
	var results []persistence.ExecutionResult
	DB.Preload("Metrics").
		Where("execution_id = ?", id).
		Find(&results)
	return results

}

func GetResultById(id uint) persistence.ExecutionResult {
	var result persistence.ExecutionResult
	DB.Preload("Result.Metrics").First(&result, id)
	return result
}

func CreateResult(result persistence.ExecutionResult) persistence.ExecutionResult {
	tx := DB.Begin()
	err := tx.Create(&result).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	return result
}

func DeleteResult(result persistence.ExecutionResult) {
	DB.Delete(&result)
}

func AddResultToExecution(execution persistence.Execution, result persistence.ExecutionResult) {
	result.ExecutionID = execution.ID
	DB.Update(result)
}
*/
