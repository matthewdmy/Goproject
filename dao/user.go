package dao

import (
	"awesomeProject/models"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (dao *UserDAO) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := dao.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDAO) Create(user *models.User) error {
	return dao.db.Create(user).Error
}
