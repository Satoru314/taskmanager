package models

import "time"

type Account struct {
	AccountID   int    `json:"account_id"`
	AccountName string `json:"account_name"`
	PathHash    string `json:"path_hash"`
}

type Task struct {
	TaskID    int       `json:"task_id"`
	AccountID int       `json:"account_id"`
	Title     string    `json:"title"`
	DueDate   time.Time `json:"due_date"`
	Progress  int       `json:"progress"`
}
