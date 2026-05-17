package models

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	DueDate   time.Time `json:"due_date"`
	Completed bool      `json:"completed"`
}

type TodoRequest struct {
	Task      string    `json:"task"`
	DueDate   time.Time `json:"due_date"`
	Completed bool      `json:"completed"`
}
