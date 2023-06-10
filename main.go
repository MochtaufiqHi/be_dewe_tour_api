package main

import (
	"dumbmerch/database"
	"os"

	// "dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	// connect database
	mysql.DatabaseInit()
	//  migration database
	database.RunMigration()

	// Route
	routes.RouteInit(e.Group("/api/v1"))
	// uploads can be used another client
	e.Static("/uploads", "./uploads")

	PORT := os.Getenv("PORT")

	fmt.Println("Server running localhost:5000")
	e.Logger.Fatal(e.Start(":" + PORT))
}
