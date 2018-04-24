package routes

import (
	"github.com/labstack/echo"
	"alexbrasser/packages/quacky-quotes/api/http"
)

func RegisterQuoteRoutes(echo *echo.Echo){
	echo.GET("/quote/:id", http.GetQuote)
	echo.GET("/quotes", http.GetQuotes)
	echo.POST("/quote", http.CreateQuote)
}
