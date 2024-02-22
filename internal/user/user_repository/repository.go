package user_repository

import (
	"github.com/google/uuid"
	"mental-atoms/internal/user/user_model"
)

var repository UserRepository

type UserRepository interface {
	Create(user user_model.User) (uuid.UUID, error)
	FindAll() ([]user_model.User, error)
}

func GetUserRepository() UserRepository {
	return repository
}

func SetUserRepository(userRepository UserRepository) {
	repository = userRepository
}
