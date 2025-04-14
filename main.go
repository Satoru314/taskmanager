package main

import (
	"log"
	"taskmanager/controllers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/", controllers.HelloHandler)

	e.GET("/tasks", controllers.GetTasksController)
	e.POST("/tasks", controllers.PostTaskController)
	e.GET("/tasks/:id", controllers.GetTaskController)
	e.PUT("/tasks/:id", controllers.PutTaskController)
	e.DELETE("/tasks/:id", controllers.DeleteTaskController)

	e.GET("/comments", controllers.GetCommentsController)
	e.POST("/comments", controllers.PostCommentController)
	e.GET("/comments/:id", controllers.GetCommentController)
	e.PUT("/comments/:id", controllers.PutCommentController)
	e.DELETE("/comments/:id", controllers.DeleteCommentController)

	e.GET("/accounts/:id", controllers.GetAccountController)
	e.POST("/accounts/:id", controllers.PostAccountController)
	e.PUT("/accounts/:id", controllers.PutAccountController)
	e.DELETE("/accounts/:id", controllers.DeleteAccountController)

	e.GET("/admin/accounts", controllers.GetAccountsController)

	log.Println("Starting server on :8080")
	log.Fatal(e.Start(":8080"))
}
