package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dev-sota/sqlboiler-example/models"
	_ "github.com/go-sql-driver/mysql" // require
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/sqlboiler_example"
	db, _ := sql.Open("mysql", dsn)
	defer db.Close()

	ctx := context.Background()

	u := &models.User{Name: "john"}
	_ = u.Insert(ctx, db, boil.Infer())
	fmt.Println(u.ID)

	get, _ := models.Users().One(ctx, db)
	fmt.Println(get.ID)

	found, _ := models.FindUser(ctx, db, u.ID)
	fmt.Println(found.ID)

	found.Name = "jane"
	rows, _ := found.Update(ctx, db, boil.Whitelist(models.UserColumns.Name))
	fmt.Println(rows)
}
