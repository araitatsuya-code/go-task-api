package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/araitatsuya-code/go-task-api/internal/model"
    "github.com/araitatsuya-code/go-task-api/internal/service"
)

// TaskHandler はタスクのHTTPハンドラー (Rails: TasksControllerに相当)
type TaskHandler struct {
    service *service.TaskService
}

// NewTaskHandler は新しいTaskHandlerを返す
func NewTaskHandler(service *service.TaskService) *TaskHandler {
    return &TaskHandler{service: service}
}

// GetTasks は全てのタスクを取得するハンドラー (Rails: index actionに相当)
func (h *TaskHandler) GetTasks(c *gin.Context) {
    tasks, err := h.service.GetAllTasks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

// GetTask は特定のタスクを取得するハンドラー (Rails: show actionに相当)
func (h *TaskHandler) GetTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    task, err := h.service.GetTaskByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    c.JSON(http.StatusOK, task)
}

// CreateTask は新しいタスクを作成するハンドラー (Rails: create actionに相当)
func (h *TaskHandler) CreateTask(c *gin.Context) {
    var task model.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.CreateTask(&task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, task)
}

// UpdateTask はタスクを更新するハンドラー (Rails: update actionに相当)
func (h *TaskHandler) UpdateTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var task model.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    task.ID = uint(id)

    if err := h.service.UpdateTask(&task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, task)
}

// DeleteTask はタスクを削除するハンドラー (Rails: destroy actionに相当)
func (h *TaskHandler) DeleteTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.service.DeleteTask(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}