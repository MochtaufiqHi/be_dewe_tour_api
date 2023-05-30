package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	UserRepository := repository.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(UserRepository)

	e.GET("/users", h.GetAllUser)
	e.GET("/user/:id", h.GetUser)
	e.POST("/user", middleware.Auth(h.AddUser))
	e.DELETE("/user/:id", middleware.Auth(h.DeleteUser))
}
