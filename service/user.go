package service

import (
	"awesomeProject/dao"
	"awesomeProject/models"
)

type UserService struct {
	userDAO *dao.UserDAO
}

func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{userDAO: userDAO}
}

func (s *UserService) Register(name, email string) (*models.User, error) {
	user := &models.User{Username: name, Email: email}
	if err := s.userDAO.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.userDAO.GetByID(id)
}
