package user_handler

import (
	"github.com/labstack/echo/v4"
	"mental-atoms/internal/user/user_service"
	"net/http"
)

func CreateUserHandler(ctx echo.Context) error {
	var createDto user_service.UserCreateDTO
	ctx.Bind(&createDto)

	userCreator := user_service.NewUserCreator()

	newUser, _ := userCreator.Create(createDto)

	return ctx.JSON(http.StatusCreated, newUser)
}
