package repositories

import (
	"bitbucket.org/instinctools/gluten/shared/persistence"
)

func init() {
	DB = persistence.InitDb()
}

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
