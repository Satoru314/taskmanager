package controllers

import "taskmanager/services"

type MyAppControllers struct {
	myAppServices services.MyAppServices
}

func NewMyAppControllers(myAppServices *services.MyAppServices) *MyAppControllers {
	return &MyAppControllers{myAppServices: *myAppServices}
}
