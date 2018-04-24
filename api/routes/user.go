package routes

import (
	"alexbrasser/api/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RegisterUserRoutes(echo *echo.Echo) {
	echo.POST("/user", http.CreateUser)
	echo.GET("/users", http.GetUsers)
	echo.GET("/user/:uuid", http.GetUser)
	echo.POST("/user/update", http.UpdateUser)

	echo.PATCH("/user/:uuid/add/permission", http.AddPermissionUser)
	echo.PATCH("/user/:uuid/delete/permission", http.DeletePermissionUser)

	echo.POST("/user/GetByEmail", http.GetUserByEmail)
	echo.POST("/user/auth", http.AuthUser)
	echo.GET("/user/self", http.SelfUser, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		ContextKey:  "user",
		TokenLookup: "header:User-Token",
	}))
	// , middleware.JWTWithConfig(middleware.JWTConfig{
	//	SigningKey:  []byte("secret"),
	//	ContextKey:  "client",
	//	TokenLookup: "header:Client-Token",
	//}))
}
