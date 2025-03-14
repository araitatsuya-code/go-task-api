package database

import (
    "fmt"
    "log"
    "os"

    "github.com/araitatsuya-code/go-task-api/internal/model"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// GetDB はデータベース接続を返す (Rails: ActiveRecord::Base.connectionに相当)
func GetDB() (*gorm.DB, error) {
    // 環境変数から設定を取得
    host := getEnv("DB_HOST", "localhost")
    user := getEnv("DB_USER", "postgres")
    password := getEnv("DB_PASSWORD", "postgres")
    dbname := getEnv("DB_NAME", "api_db")
    port := getEnv("DB_PORT", "5432")
    sslmode := getEnv("DB_SSLMODE", "disable")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        host, user, password, dbname, port, sslmode)
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}

// SetupDB はマイグレーションを実行 (Rails: db:migrateに相当)
func SetupDB() (*gorm.DB, error) {
    db, err := GetDB()
    if err != nil {
        return nil, err
    }

    // マイグレーション実行
    err = db.AutoMigrate(&model.Task{})
    if err != nil {
        return nil, err
    }

    log.Println("Database migration completed")
    return db, nil
}

// 環境変数取得ヘルパー
func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}