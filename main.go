package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Kazzzz195/GoProject/controllers"
	"github.com/Kazzzz195/GoProject/routers"
	"github.com/Kazzzz195/GoProject/services"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}
	ser := services.NewTodoService(db)
	con := controllers.NewTodoController(ser)

	r := routers.NewRouter(con)
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
