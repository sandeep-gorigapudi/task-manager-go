package model

type Task struct {
	ID          int 	`json:"id"`
	Description string 	`json:"description" binding:"required"`
	Completed   bool	`json:"completed"`
}