package service

import (
	"go-task-api/internal/database"
	"go-task-api/internal/model"
)

type TaskManagerInterface interface {
	Create(desc string) error
	ReadAll() ([]*model.Task, error)
	ReadOne(id int) (*model.Task, error)
	UpdateTask(id int, desc string) error
	DeleteTask(id int) error
	MarkDone(id int) error
}

type TaskManager struct{}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (tm *TaskManager) Create(desc string) error {
	query := "INSERT INTO tasks (description) VALUES (?)"
	_, err := db.DB.Exec(query, desc)
	return err
}


func (tm *TaskManager) ReadAll() ([]*model.Task, error) {
	rows, err := db.DB.Query("SELECT id, description, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Description, &task.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}



func (tm *TaskManager) ReadOne(id int) (*model.Task, error) {
	query := "SELECT id, description, completed FROM tasks WHERE id = ?"

	row := db.DB.QueryRow(query, id)
	var task model.Task

	err := row.Scan(&task.ID, &task.Description, &task.Completed)
	if err != nil {
		return nil, err
	}

	return &task, nil
}


func (tm *TaskManager) UpdateTask(id int, desc string) error {
	query := "UPDATE tasks SET description = ? WHERE id = ?"
	_, err := db.DB.Exec(query, desc, id)
	return err
}


func (tm *TaskManager) DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = ?"
	_, err := db.DB.Exec(query, id)
	return err
}



func (tm *TaskManager) MarkDone(id int) error {
	query := "UPDATE tasks SET completed = true WHERE id = ?"
	_, err := db.DB.Exec(query, id)
	return err
}