package bkp


import (
    "go-task-api/internal/model"
)

type TaskManagerInterface interface {
	Create(desc string)
	ReadAll() []*model.Task
	ReadOne(id int) *model.Task
	UpdateTask(id int, desc string) bool
	DeleteTask(id int) bool
	MarkDone(id int) bool
}

type TaskManager struct {
	tasks map[int]*model.Task
	order []int
	nextID  int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  make(map[int]*model.Task),
		order: []int{},
		nextID: 1,
	}
}

func (tm *TaskManager) Create(desc string) {
	task := &model.Task{
		ID : tm.nextID,
		Description : desc,
		Completed : false,
	}
	tm.tasks[tm.nextID] = task
	tm.order = append(tm.order, tm.nextID)
	tm.nextID++
}

func (tm *TaskManager) ReadAll() []*model.Task {
	 allTasks := make([]*model.Task, 0, len(tm.order))
	for _, val := range tm.order {
		allTasks = append(allTasks, tm.tasks[val])
	}
	return allTasks
}

func (tm *TaskManager) ReadOne(id int) *model.Task {
	if task, exists := tm.tasks[id]; exists{
		return task
	}
	return nil
}

func (tm *TaskManager) UpdateTask(id int, desc string) bool {
	if t, exists := tm.tasks[id]; exists {
		t.Description = desc
		return true
	}
	return false
}


func (tm *TaskManager) DeleteTask(id int) bool{
	if _, exists := tm.tasks[id]; exists {
		delete(tm.tasks, id)
		for i, taskID := range tm.order {
			if taskID == id {
				tm.order = append(tm.order[:i], tm.order[i+1:]...)
				return true
			}
		}
	}
	return false
}

func (tm *TaskManager) MarkDone(id int) bool{
	if t, exists := tm.tasks[id]; exists {
		t.Completed = true
		return true	
	}
	return false
}
