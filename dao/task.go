package dao

import (
	"awesomeProject/models"

	"gorm.io/gorm"
)

type TaskDAO struct {
	db *gorm.DB
}

func (d *TaskDAO) Create(task *models.Task) error {
	return d.db.Create(task).Error
}

func NewTaskDAO(db *gorm.DB) *TaskDAO {
	return &TaskDAO{db: db}
}

func (d *TaskDAO) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	if err := d.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (d *TaskDAO) DeleteByID(id uint) error {
	return d.db.Delete(&models.Task{}, id).Error
}

func (d *TaskDAO) Update(task *models.Task) error {
	return d.db.Save(task).Error
}
