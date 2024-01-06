package service

import (
	"user-management/models"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type IUserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	CreateUser(user models.User) error
	UpdateUserByID(id uint, user models.User) error
	DeleteByID(id uint) error
}

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) IUserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		log.Error("Error while fetching all users: ", err)
		return nil, err
	}
	log.Info("All users fetched successfully Users: ", users)
	return users, nil
}

func (s *UserService) GetUserByID(id uint) (models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		log.Error("Error while fetching user by id: ", err)
		return user, err
	}
	log.Info("User fetched successfully User: ", user)
	return user, nil
}

func (s *UserService) CreateUser(user models.User) error {
	if err := s.db.Create(&user).Error; err != nil {
		log.Error("Error while creating user: ", err)
		return err
	}
	log.Info("User created successfully User: ", user)
	return nil
}

func (s *UserService) UpdateUserByID(id uint, user models.User) error {
	if err := s.db.Model(&user).Where("id = ?", id).Updates(models.User{Name: user.Name, Age: user.Age}).Error; err != nil {
		log.Error("Error while updating user by id: ", err)
		return err
	}
	log.Info("User updated successfully User: ", user)
	return nil
}

func (s *UserService) DeleteByID(id uint) error {
	if err := s.db.Delete(&models.User{}, id).Error; err != nil {
		log.Error("Error while deleting user by id: ", err)
		return err
	}
	log.Info("User deleted successfully User ID: ", id)
	return nil
}
