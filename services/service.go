package services

import "database/sql"

type MyAppServices struct {
	db *sql.DB
}

func NewMyAppServices(db *sql.DB) *MyAppServices {
	return &MyAppServices{db: db}
}
