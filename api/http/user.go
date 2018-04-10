package http

import (
	"github.com/labstack/echo"
	"alexbrasser/model"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"strconv"
)

func GetUser(context echo.Context) (error) {
	user := &model.User{}

	if err := model.GetUser(user, context.Param("id")); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, &user)
}

func GetUsers(context echo.Context) (error) {
	users := &[]model.User{}
	if err := model.GetUsers(users); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, &users)
}

func CreateUser(context echo.Context) (error) {
	user := &model.User{}

	if err := context.Bind(&user); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	if err := model.CreateUser(user); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, &user)
}

func AuthenticateUser(context echo.Context) (error) {
	user := &model.User{}

	if err := context.Bind(&user); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	model.GetUserByEmailAndPassword(user)

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = &user.ID
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

func SelfUser(context echo.Context) (error){
	user := &model.User{}
	id := context.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["id"].(float64)
	model.GetUser(user, strconv.Itoa(int(id)))
	return context.JSON(http.StatusOK, &user)
}
