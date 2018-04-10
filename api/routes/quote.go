package routes

import (
	"github.com/labstack/echo"
	"alexbrasser/api/http"
)

func RegisterQuoteRoutes(echo *echo.Echo){
	echo.GET("/quotes", http.GetQuotes)
	echo.POST("/quote", http.CreateQuote)
}
