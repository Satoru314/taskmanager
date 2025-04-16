package main

import (
	"database/sql"
	"fmt"
	"log"
	"taskmanager/controllers"
	"taskmanager/services"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var (
	connStr = fmt.Sprintf("satoru:satoru@localhost:5432/taskmanager?sslmode=disable")
)

func main() {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	ser := services.NewMyAppServices(db)
	con := controllers.NewMyAppControllers(ser)
	defer db.Close()

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
