package http

import (
	"github.com/labstack/echo"
	"alexbrasser/model"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"strconv"
)

func GetClient(context echo.Context) (error) {
	client := &model.Client{}

	if err := model.GetClient(client, context.Param("id")); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, &client)
}


func CreateClient(context echo.Context) (error) {
	client := &model.Client{}

	if err := context.Bind(&client); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	if err := model.CreateClient(client); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, &client)
}

func AuthenticateClient(context echo.Context) (error) {
	client := &model.Client{}

	if err := context.Bind(&client); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	model.GetClientBySecret(client)

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = &client.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("VZHE7JAPWMUI8KFHC6Z020TV9P2J8N1KIU86ZKGVCSJ1RFMRXH87MXX6H14TC0VA"))
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func SelfClient(context echo.Context) (error){
	client := &model.Client{}
	id := context.Get("client").(*jwt.Token).Claims.(jwt.MapClaims)["id"].(float64)
	model.GetClient(client, strconv.Itoa(int(id)))
	return context.JSON(http.StatusOK, &client)
}