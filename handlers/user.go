package handlers

import (
	"context"
	"net/http"

	"github.com/dev-sota/sqlboiler-example/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/boil"
)

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	u := &models.User{Name: name}
	_ = u.Insert(context.Background(), db, boil.Infer())
	return c.JSON(http.StatusOK, u)
}
