package http

import (
	"github.com/labstack/echo"
	"net/http"
	"alexbrasser/packages/quacky-quotes/model"
)

func GetQuote(context echo.Context) (error) {
	quote := &model.Quote{}

	if err := quote.Get(context.Param("id")); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, &quote)
}

func GetQuotes(context echo.Context) (error) {
	quotes := &model.Quotes{}
	if err := quotes.Get(); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, &quotes)
}

func CreateQuote(context echo.Context) (error) {
	quote := &model.Quote{}

	if err := context.Bind(&quote); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	if err := quote.Create(); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, &quote)
}
