package repository

import "practice02/internal/domain"

type TodoRepository interface {
    Create(todo *domain.Todo) error
    GetAll() ([]domain.Todo, error)
    GetByID(id uint) (*domain.Todo, error)
    Update(todo *domain.Todo) error
    Delete(id uint) error
}