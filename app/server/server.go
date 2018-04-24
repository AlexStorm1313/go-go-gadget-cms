package server

import (
	"github.com/labstack/echo"
	"alexbrasser/api/routes"
	"github.com/labstack/echo/middleware"
	quackyquotesRoutes "alexbrasser/packages/quacky-quotes/api/routes"
	"alexbrasser/model"
	"alexbrasser/app/cache"
)

func Run() {
	echo := echo.New()
	echo.Debug = true
	echo.Use(middleware.Logger())
	echo.Use(middleware.CORS())
	routes.RegisterUserRoutes(echo)
	quackyquotesRoutes.RegisterQuoteRoutes(echo)
	routes.RegisterClientRoutes(echo)
	routes.RegisterActionRoutes(echo)

	client := cache.OpenRedis()
	client.FlushDB()
	client.Close()

	outes := echo.Routes()
	for i := 0; i < len(outes); i++ {
		oute := &model.Action{Method: outes[i].Method, Path: outes[i].Path, Name: outes[i].Name}
		oute.Create(*oute)
	}

	echo.Logger.Fatal(echo.Start(":3000"))
}
