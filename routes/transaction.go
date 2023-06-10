package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	TransactionRepository := repository.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(TransactionRepository)

	e.GET("/transaction", h.GetAllTransaction)
	e.GET("/transaction/:id", h.GetTransaction)
	e.GET("/transaction-user/:id", h.GetTransactionByUser)
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	// e.PATCH("/transaction/:id", middleware.Auth(h.UpdateTransaction))
	e.POST("/notification", h.Notification)
	e.DELETE("/transaction/:id", h.DeleteTransaction)
}
