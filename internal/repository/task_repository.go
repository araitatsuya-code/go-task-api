package repository

import (
    "github.com/araitatsuya-code/go-task-api/internal/model"
    "gorm.io/gorm"
)

// TaskRepository はタスクのデータアクセス層 (Rails: TaskモデルのActiveRecordメソッドに相当)
type TaskRepository struct {
    db *gorm.DB
}

// NewTaskRepository は新しいTaskRepositoryを返す
func NewTaskRepository(db *gorm.DB) *TaskRepository {
    return &TaskRepository{db: db}
}

// FindAll は全てのタスクを取得 (Rails: Task.allに相当)
func (r *TaskRepository) FindAll() ([]model.Task, error) {
    var tasks []model.Task
    result := r.db.Find(&tasks)
    return tasks, result.Error
}

// FindByID はIDによるタスク取得 (Rails: Task.find(id)に相当)
func (r *TaskRepository) FindByID(id uint) (model.Task, error) {
    var task model.Task
    result := r.db.First(&task, id)
    return task, result.Error
}

// Create は新しいタスクを作成 (Rails: Task.create!に相当)
func (r *TaskRepository) Create(task *model.Task) error {
    return r.db.Create(task).Error
}

// Update はタスクを更新 (Rails: task.update!に相当)
func (r *TaskRepository) Update(task *model.Task) error {
    return r.db.Save(task).Error
}

// Delete はタスクを削除 (Rails: task.destroy!に相当)
func (r *TaskRepository) Delete(id uint) error {
    return r.db.Delete(&model.Task{}, id).Error
}