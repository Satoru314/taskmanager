package main

import (
	"log"
	"taskmanager/controllers"

	"github.com/labstack/echo/v4"
)

func main() {

	c := echo.New()

	c.GET("/", controllers.HelloHandler)

	c.GET("/tasks/:account_id", controllers.GetTasksHandler)
	c.POST("/tasks/:account_id", controllers.PostTaskHandler)
	c.PUT("/tasks/:task_id", controllers.PutTaskHandler)
	c.DELETE("/tasks/:task_id", controllers.DeleteTaskHandler)

	c.GET("/accounts/:account_id", controllers.GetAccountHandler)
	c.POST("/accounts", controllers.PostAccountHandler)
	c.PUT("/accounts/:account_id", controllers.PutAccountHandler)
	c.DELETE("/accounts/:account_id", controllers.DeleteAccountHandler)

	log.Println("Starting server on :8080")
	log.Fatal(c.Start(":8080"))
}
