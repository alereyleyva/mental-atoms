package user_service

import (
	"mental-atoms/internal/user/user_model"
	"mental-atoms/internal/user/user_repository"
)

type UserFinder struct {
	repository user_repository.UserRepository
}

func (finder UserFinder) FindAll() ([]user_model.User, error) {
	return finder.repository.FindAll()
}

func NewUserFinder() UserFinder {
	return UserFinder{repository: user_repository.GetUserRepository()}
}
