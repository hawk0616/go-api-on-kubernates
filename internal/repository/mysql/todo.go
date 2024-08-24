package mysql

import (
    "gorm.io/gorm"
    "practice02/internal/domain"
    "practice02/internal/domain/repository"
)

type todoRepositoryMysql struct {
    db *gorm.DB
}

func NewTodoRepositoryMysql(db *gorm.DB) repository.TodoRepository {
    return &todoRepositoryMysql{db}
}

func (r *todoRepositoryMysql) Create(todo *domain.Todo) error {
    return r.db.Create(todo).Error
}

func (r *todoRepositoryMysql) GetAll() ([]domain.Todo, error) {
    var todos []domain.Todo
    err := r.db.Find(&todos).Error
    return todos, err
}

func (r *todoRepositoryMysql) GetByID(id uint) (*domain.Todo, error) {
    var todo domain.Todo
    err := r.db.First(&todo, id).Error
    return &todo, err
}

func (r *todoRepositoryMysql) Update(todo *domain.Todo) error {
    return r.db.Save(todo).Error
}

func (r *todoRepositoryMysql) Delete(id uint) error {
    return r.db.Delete(&domain.Todo{}, id).Error
}