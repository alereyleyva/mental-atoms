package user_repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"mental-atoms/internal/user/user_model"
	"mental-atoms/pkg/database"
)

type GormUserRepository struct {
	connection *gorm.DB
}

func (repository GormUserRepository) FindAll() ([]user_model.User, error) {
	var users []user_model.User

	repository.connection.Find(&users)

	return users, nil
}

func (repository GormUserRepository) Create(user user_model.User) (uuid.UUID, error) {
	repository.connection.Create(&user)

	return user.ID, nil
}

func NewGormUserRepository() GormUserRepository {
	gormRepository := GormUserRepository{connection: database.GetConnectionInstance()}

	err := gormRepository.connection.AutoMigrate(&user_model.User{})

	if nil != err {
		log.Fatal(err)
	}

	return gormRepository
}
