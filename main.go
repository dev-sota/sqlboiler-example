package main

import (
	"database/sql"

	"github.com/dev-sota/sqlboiler-example/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql" // require
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/sqlboiler_example"
	db, _ := sql.Open("mysql", dsn)
	defer db.Close()
	handlers.SetDB(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", handlers.CreateUser)

	e.Logger.Fatal(e.Start("localhost:1323"))
}
