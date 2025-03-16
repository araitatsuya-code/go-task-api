package service

import (
    "errors"

    "github.com/araitatsuya-code/go-task-api/internal/model"
    "github.com/araitatsuya-code/go-task-api/internal/repository"
)

// TaskService はタスクのビジネスロジック層 (Rails: TaskServiceクラスに相当)
type TaskService struct {
    repo *repository.TaskRepository
}

// NewTaskService は新しいTaskServiceを返す
func NewTaskService(repo *repository.TaskRepository) *TaskService {
    return &TaskService{repo: repo}
}

// GetAllTasks は全てのタスクを取得
func (s *TaskService) GetAllTasks() ([]model.Task, error) {
    return s.repo.FindAll()
}

// GetTaskByID はIDによるタスク取得
func (s *TaskService) GetTaskByID(id uint) (model.Task, error) {
    return s.repo.FindByID(id)
}

// CreateTask は新しいタスクを作成
func (s *TaskService) CreateTask(task *model.Task) error {
    // ここに必要なバリデーションなどを追加
    if task.Title == "" {
        return errors.New("title cannot be empty")
    }
    return s.repo.Create(task)
}

// UpdateTask はタスクを更新
func (s *TaskService) UpdateTask(task *model.Task) error {
    // タスクの存在確認
    _, err := s.repo.FindByID(task.ID)
    if err != nil {
        return err
    }
    return s.repo.Update(task)
}

// DeleteTask はタスクを削除
func (s *TaskService) DeleteTask(id uint) error {
    // タスクの存在確認
    _, err := s.repo.FindByID(id)
    if err != nil {
        return err
    }
    return s.repo.Delete(id)
}