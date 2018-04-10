package routes

import (
	"github.com/labstack/echo"
	"alexbrasser/api/http"
	"github.com/labstack/echo/middleware"
)

func RegisterClientRoutes(echo *echo.Echo) {
	echo.POST("/clients", http.CreateClient)
	echo.GET("/client/:id", http.GetClient)
	echo.POST("/client/authenticate", http.AuthenticateClient)
	echo.GET("/client/self", http.SelfClient, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("VZHE7JAPWMUI8KFHC6Z020TV9P2J8N1KIU86ZKGVCSJ1RFMRXH87MXX6H14TC0VA"),
		ContextKey:  "client",
		TokenLookup: "header:Client-Token",
	}))
}
