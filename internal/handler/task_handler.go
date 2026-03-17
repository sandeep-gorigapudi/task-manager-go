package handler

import (
    "net/http"
	"strconv"
    "go-task-api/internal/service"
	"go-task-api/internal/model"
    "github.com/gin-gonic/gin"
)

type TaskHandler struct {
    taskService  service.TaskManager
}

func NewTaskHandler(svc service.TaskManager) *TaskHandler {
    return &TaskHandler{taskService : svc}
}

// create task
func (h *TaskHandler) CreateTask(c *gin.Context) {
    var req model.Task
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    h.taskService.Create(req.Description)
    c.JSON(http.StatusCreated, gin.H{"message": "task created"})
}

//Read all tasks
func (h *TaskHandler) ListTasks(c *gin.Context) {
    tasks, err := h.taskService.ReadAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

//Read one task by id
func (h *TaskHandler) GetTask(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
        return
    }

    task, err := h.taskService.ReadOne(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}

//Update task description by id
func (h *TaskHandler) UpdateTask(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
        return
    }

    var req model.Task
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.taskService.UpdateTask(id, req.Description); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "task updated"})
}


//Delete task by id
func (h *TaskHandler) DeleteTask(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
        return
    }

    if err := h.taskService.DeleteTask(id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}

//Mark task as done by id
func (h *TaskHandler) MarkDone(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
        return
    }

    if err := h.taskService.MarkDone(id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "task marked done"})
}