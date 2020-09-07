package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/dev-sota/sqlboiler-example/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/boil"
)

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	u := &models.User{Name: name}
	_ = u.Insert(context.Background(), db, boil.Infer())
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	u, _ := models.FindUser(context.Background(), db, id)
	return c.JSON(http.StatusOK, u)
}
