package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func CountryRoutes(e *echo.Group) {
	CountryRepository := repository.RepositoryCountry(mysql.DB)
	h := handlers.HandlerCountry(CountryRepository)

	e.GET("/country", h.GetAllCountry)
	e.GET("/country/:id", h.GetCountry)
	e.POST("/country", middleware.Auth(h.CreateCountry))
	e.PATCH("/country/:id", middleware.Auth(h.UpdateCountry))
	e.DELETE("/country/:id", middleware.Auth(h.DeleteCountry))
}
