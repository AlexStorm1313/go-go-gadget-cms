package routes

import (
	"github.com/labstack/echo"
	"alexbrasser/api/http"
)

func RegisterActionRoutes(echo *echo.Echo) {
	echo.GET("/actions", http.GetActions)
	echo.GET("/action/:uuid", http.GetAction)
	echo.POST("/action/:uuid/update", http.UpdateAction)
	echo.PATCH("/action/:uuid/add/permission", http.AddPermissionAction)
	echo.PATCH("/action/:uuid/delete/permission", http.DeletePermissionAction)
	echo.GET("/action/:uuid/get/permissions", http.GetPermissionsAction)
}
