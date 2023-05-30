package main

import (
	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	e := echo.New()

	// connect database
	mysql.DatabaseInit()
	//  migration database
	database.RunMigration()

	// Route
	routes.RouteInit(e.Group("/api/v1"))
	// uploads can be used another client
	e.Static("/uploads", "./uploads")

	fmt.Println("Server running localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
