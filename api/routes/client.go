package routes

import (
	"github.com/labstack/echo"
	"alexbrasser/api/http"
	"github.com/labstack/echo/middleware"
)

func RegisterClientRoutes(echo *echo.Echo) {
	echo.POST("/client", http.CreateClient)
	echo.GET("/client/:uuid", http.GetClient)
	echo.POST("/client/auth", http.AuthClient)
	echo.GET("/client/self", http.SelfClient, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		ContextKey:  "client",
		TokenLookup: "header:Client-Token",
	}))
	echo.PATCH("/client/:uuid/add/permission", http.AddPermissionClient)
	echo.PATCH("/client/:uuid/delete/permission", http.DeletePermissionClient)
}
