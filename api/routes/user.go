package routes

import (
	"alexbrasser/api/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RegisterUserRoutes(echo *echo.Echo) {
	echo.POST("/user", http.CreateUser)
	echo.GET("/users", http.GetUsers)
	echo.GET("/user/:id", http.GetUser)
	echo.POST("/user/authenticate", http.AuthenticateUser)
	echo.GET("/user/self", http.SelfUser, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("VZHE7JAPWMUI8KFHC6Z020TV9P2J8N1KIU86ZKGVCSJ1RFMRXH87MXX6H14TC0VA"),
		ContextKey:  "user",
		TokenLookup: "header:User-Token",
	}), middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("VZHE7JAPWMUI8KFHC6Z020TV9P2J8N1KIU86ZKGVCSJ1RFMRXH87MXX6H14TC0VA"),
		ContextKey:  "client",
		TokenLookup: "header:Client-Token",
	}))
}
