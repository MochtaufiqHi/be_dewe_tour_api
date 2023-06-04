package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func TripRoutes(e *echo.Group) {
	TripRepository := repository.RepositoryTrip(mysql.DB)
	h := handlers.HandlerTrip(TripRepository)

	e.GET("/trip", h.GetAllTrip)
	e.GET("/trip/:id", h.GetTrip)
	// e.POST("/trip", h.CreateTrip)
	e.POST("/trip", middleware.Auth(middleware.UploadFile(h.CreateTrip)))
	e.PATCH("/trip/:id", middleware.Auth(middleware.UploadFile(h.UpdateTrip)))
	e.DELETE("/trip/:id", middleware.Auth(h.DeleteTrip))
}
