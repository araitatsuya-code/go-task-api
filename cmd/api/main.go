package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/araitatsuya-code/go-task-api/internal/handler"
    "github.com/araitatsuya-code/go-task-api/internal/repository"
    "github.com/araitatsuya-code/go-task-api/internal/service"
    "github.com/araitatsuya-code/go-task-api/pkg/database"
)

func main() {
    // データベース接続 (Rails: database.ymlと接続に相当)
    db, err := database.SetupDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // リポジトリ、サービス、ハンドラーの初期化
    taskRepo := repository.NewTaskRepository(db)
    taskService := service.NewTaskService(taskRepo)
    taskHandler := handler.NewTaskHandler(taskService)

    // Ginルーターの設定 (Rails: routes.rbに相当)
    r := gin.Default()

    // ヘルスチェックエンドポイント
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "ok",
        })
    })

    // APIルーティング (Rails: resources :tasksに相当)
    api := r.Group("/api/v1")
    {
        tasks := api.Group("/tasks")
        {
            tasks.GET("", taskHandler.GetTasks)         // GET /api/v1/tasks - index
            tasks.GET("/:id", taskHandler.GetTask)      // GET /api/v1/tasks/:id - show
            tasks.POST("", taskHandler.CreateTask)      // POST /api/v1/tasks - create
            tasks.PUT("/:id", taskHandler.UpdateTask)   // PUT /api/v1/tasks/:id - update
            tasks.DELETE("/:id", taskHandler.DeleteTask) // DELETE /api/v1/tasks/:id - destroy
        }
    }

    // サーバー起動 (Rails: rails s -p 8080に相当)
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}