package http

import (
	"github.com/labstack/echo"
	"alexbrasser/model"
	"net/http"
)

func GetQuote(context echo.Context) (error) {
	quote := &model.Quote{}

	if err := model.GetQuote(quote, context.Param("id")); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, &quote)
}

func GetQuotes(context echo.Context) (error) {
	quotes := &[]model.Quote{}
	if err := model.GetQuotes(quotes); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, &quotes)
}

func CreateQuote(context echo.Context) (error) {
	quote := &model.Quote{}

	if err := context.Bind(&quote); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	if err := model.CreateQuote(quote); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, &quote)
}
