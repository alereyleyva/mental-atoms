package user_router

import (
	"github.com/labstack/echo/v4"
	"mental-atoms/internal/user/user_handler"
)

func SetUserRoutes(g *echo.Group) {
	userGroup := g.Group("/users")

	userGroup.GET("", user_handler.FindUsersHandler)
	userGroup.POST("", user_handler.CreateUserHandler)
}
