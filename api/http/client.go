package http

import (
	"github.com/labstack/echo"
	"alexbrasser/model"
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

func GetClient(context echo.Context) (error) {
	client := &model.Client{}

	if err := client.Get(context.Param("id")); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, &client)
}

func CreateClient(context echo.Context) (error) {
	client := &model.Client{}
	data := &model.Client{}

	if err := context.Bind(&data); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	if err := client.Create(*data); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, &client)
}

func AuthClient(context echo.Context) (error) {
	client := &model.Client{}

	if err := context.Bind(&client); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	if err := client.GetBySecret(client.Secret); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	err, token := client.CreateToken()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, token)
}

func AddPermissionClient(context echo.Context) (error) {
	client := &model.Client{}
	permission := &model.Permission{}
	if err := client.GetByUUID(context.Param("uuid")); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := context.Bind(&permission); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := client.AddPermission(*permission); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, &client)
}

func DeletePermissionClient(context echo.Context)(error){
	client := &model.Client{}
	permission := &model.Permission{}
	if err := client.GetByUUID(context.Param("uuid")); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := context.Bind(&permission); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := client.DeletePermission(*permission); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, &client)
}

func SelfClient(context echo.Context) (error){
	client := &model.Client{}
	uuid := context.Get("client").(*jwt.Token).Claims.(jwt.MapClaims)["client"].(string)
	client.GetByUUID(uuid)
	return context.JSON(http.StatusOK, &client)
}
