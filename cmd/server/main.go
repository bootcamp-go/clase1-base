package main

import (
	"database/sql"

	"github.com/bootcamp-go/clase1-base/cmd/server/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := gin.Default()

	db, err := sql.Open("mysql", "root:root@tcp()")
	if err != nil {
		panic(err)
	}

	router.NewRouter(r, db).MapRoutes()

}
