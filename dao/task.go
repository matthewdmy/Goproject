package dao

import (
	"awesomeProject/models"

	"gorm.io/gorm"
)

type TaskDAO struct {
	db *gorm.DB
}

func (d TaskDAO) Create(task *models.Task) error {
	return d.db.Create(task).Error
}

func NewTaskDAO(db *gorm.DB) *TaskDAO {
	return &TaskDAO{db: db}
}
