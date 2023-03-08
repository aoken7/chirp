package main

import (
	"chirp/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routers.Setup(e)

	e.Logger.Fatal(e.Start(":80"))
}
