package model

import (
    "time"
)

// Task はタスクのモデル (Rails: Task modelに相当)
type Task struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Title       string    `json:"title" binding:"required"`
    Description string    `json:"description"`
    Status      string    `json:"status" gorm:"default:pending"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}