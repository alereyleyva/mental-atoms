package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"mental-atoms/internal/user/user_repository"
	"mental-atoms/internal/user/user_router"
	"mental-atoms/pkg/database"
	"net/http"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initializeEntityRepositories() {
	user_repository.SetUserRepository(user_repository.NewGormUserRepository())
}

func setApiRoutes(e *echo.Echo) {
	apiGroup := e.Group("/api")

	apiGroup.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Mental Atoms API")
	})

	user_router.SetUserRoutes(apiGroup)
}

func main() {
	e := echo.New()

	loadEnv()

	err := database.ConnectDatabase()

	initializeEntityRepositories()

	if nil != err {
		log.Fatal(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Mental Atoms")
	})

	setApiRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
