package server

import (
	"github.com/labstack/echo"
	"alexbrasser/api/routes"
	"github.com/labstack/echo/middleware"
)

func Run() {
	echo := echo.New()
	echo.Debug = true
	echo.Use(middleware.Logger())
	echo.Use(middleware.CORS())
	routes.RegisterUserRoutes(echo)
	routes.RegisterQuoteRoutes(echo)
	routes.RegisterClientRoutes(echo)
	echo.Logger.Fatal(echo.Start(":3000"))
}
