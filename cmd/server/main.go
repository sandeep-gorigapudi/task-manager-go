package main

import (
    "go-task-api/internal/handler"
    "go-task-api/internal/service"
    "github.com/gin-gonic/gin"
    "go-task-api/internal/database"
)

func main() {
    // Create service instance
    db.Connect()
    taskSvc := service.NewTaskManager()

    // Create handler with the service
    taskHandler := handler.NewTaskHandler(*taskSvc)

    r := gin.Default()

    // Define routes
    r.POST("/tasks", taskHandler.CreateTask)
    r.GET("/tasks", taskHandler.ListTasks)
    r.GET("/tasks/:id", taskHandler.GetTask)
    r.PUT("/tasks/:id", taskHandler.UpdateTask)
    r.DELETE("/tasks/:id", taskHandler.DeleteTask)
    r.PATCH("/tasks/:id/done", taskHandler.MarkDone)

    r.Run(":8080")
}