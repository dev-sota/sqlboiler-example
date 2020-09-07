package handlers

import (
	"context"
	"net/http"

	"github.com/dev-sota/sqlboiler-example/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/boil"
)

func CreateUser(c echo.Context) error {
	u := &models.User{Name: "john"}
	_ = u.Insert(context.Background(), db, boil.Infer())
	return c.String(http.StatusOK, u.Name)
}
