package http

import (
	"alexbrasser/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetActions(context echo.Context) error {
	actions := &model.Actions{}

	if err := actions.Get(); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, &actions)
}

func GetAction(context echo.Context) error {
	action := &model.Action{}

	if err := action.GetByUUID(context.Param("uuid")); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, &action)
}

func UpdateAction(context echo.Context) error {
	action := &model.Action{}
	data := &model.Action{}

	if err := action.GetByUUID(context.Param("uuid")); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := context.Bind(&data); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := action.Update(*data); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, &action)
}

func GetPermissionsAction(context echo.Context) error {
	action := &model.Action{}

	if err := action.GetByUUID(context.Param("uuid")); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	if err := action.GetPermissions(); err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, &action)
}

func AddPermissionAction(context echo.Context) error {
	action := &model.Action{}
	permission := &model.Permission{}
	if err := action.GetByUUID(context.Param("uuid")); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := context.Bind(&permission); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := action.AddPermission(*permission); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}


	return context.JSON(http.StatusOK, &action)
}

func DeletePermissionAction(context echo.Context) error {
	action := &model.Action{}
	permission := &model.Permission{}
	if err := action.GetByUUID(context.Param("uuid")); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := context.Bind(&permission); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := action.DeletePermission(*permission); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, &action)
}
