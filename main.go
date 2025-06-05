package main

import (
	"database/sql"
	"log"
	"taskmanager/controllers"
	"taskmanager/services"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var (
	connStr = ("host=127.0.0.1 port=5432 user=docker password=docker dbname=task_manager sslmode=disable")
)

func main() {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	ser := services.NewMyAppServices(db)
	con := controllers.NewMyAppControllers(ser)

	c := echo.New()

	c.GET("/", con.HelloHandler)

	c.GET("/tasks/:account_id", con.GetTasksHandler)
	c.POST("/tasks/:account_id", con.PostTaskHandler)
	c.PUT("/tasks/:task_id", con.PutTaskHandler)
	c.DELETE("/tasks/:task_id", con.DeleteTaskHandler)

	c.GET("/accounts/:account_id", con.GetAccountHandler)
	c.POST("/accounts", con.PostAccountHandler)
	c.PUT("/accounts/:account_id", con.PutAccountHandler)
	c.DELETE("/accounts/:account_id", con.DeleteAccountHandler)

	log.Println("Starting server on :8080")
	log.Fatal(c.Start(":8080"))
}
