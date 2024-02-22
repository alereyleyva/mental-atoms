package user_handler

import (
	"github.com/labstack/echo/v4"
	"mental-atoms/internal/user/user_service"
	"net/http"
)

func FindUsersHandler(ctx echo.Context) error {
	userFinder := user_service.NewUserFinder()

	users, _ := userFinder.FindAll()

	return ctx.JSON(http.StatusOK, users)
}
