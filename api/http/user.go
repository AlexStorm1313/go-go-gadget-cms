package http

import (
	"alexbrasser/model"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetUser(context echo.Context) error {
	user := &model.User{}

	if err := user.GetByUUID(context.Param("uuid")); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, &user)
}

func GetUsers(context echo.Context) error {
	users := &model.Users{}

	if err := users.Get(); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, &users)
}

func CreateUser(context echo.Context) error {
	user := &model.User{}
	data := &model.User{}

	if err := context.Bind(&data); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	if err := user.Create(*data); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, &user)
}

func UpdateUser(context echo.Context) error {
	user := &model.User{}
	data := &model.User{}

	if err := context.Bind(&data); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	if err := user.Update(*data); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, &user)
}

func GetUserByEmail(context echo.Context) error {
	user := &model.User{}

	if err := context.Bind(&user); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	if err := user.GetByEmail(user.Email); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, &user)

}

func AuthUser(context echo.Context) error {
	user := &model.User{}

	if err := context.Bind(&user); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	password := user.Password

	if err := user.GetByEmail(user.Email); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := user.CheckPasswordHash(password); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	err, token := user.CreateToken()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, token)

}

func SelfUser(context echo.Context) error {
	user := &model.User{}
	uuid := context.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["user"].(string)
	user.GetByUUID(uuid)
	return context.JSON(http.StatusOK, &user)
}

func AddPermissionUser(context echo.Context) error {
	user := &model.User{}
	permission := &model.Permission{}
	if err := user.GetByUUID(context.Param("uuid")); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := context.Bind(&permission); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := user.AddPermission(*permission); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, &user)
}

func DeletePermissionUser(context echo.Context) error {
	user := &model.User{}
	permission := &model.Permission{}
	if err := user.GetByUUID(context.Param("uuid")); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := context.Bind(&permission); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := user.DeletePermission(*permission); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, &user)
}
