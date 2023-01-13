package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"sample-api-rest/routes"
	"sample-api-rest/storage"
)

func main() {
	e := echo.New()
	storage.NewDB()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Customer(e)

	e.Logger.Fatal(e.Start(":8080"))
}
