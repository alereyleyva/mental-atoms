package user_service

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mental-atoms/internal/user/user_model"
	"mental-atoms/internal/user/user_repository"
	"time"
)

type UserCreateDTO struct {
	Email    string
	Password string
}

type UserCreator struct {
	repository user_repository.UserRepository
}

func (creator UserCreator) Create(createDto UserCreateDTO) (*user_model.User, error) {
	passwordBytes := []byte(createDto.Password)

	hash, err := bcrypt.GenerateFromPassword(passwordBytes, 0)

	if nil != err {
		return nil, fmt.Errorf("error hashing atom password: %s", err)
	}

	newUser := user_model.User{
		ID:        uuid.New(),
		Email:     createDto.Email,
		Password:  string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, _ := creator.repository.Create(newUser)

	log.Printf("created router: %s", id)

	return &newUser, nil
}

func NewUserCreator() UserCreator {
	return UserCreator{repository: user_repository.GetUserRepository()}
}
