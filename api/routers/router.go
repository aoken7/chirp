package routers

import (
	"chirp/controllers"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	e.GET("/", controllers.Hello)
}
