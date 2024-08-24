package usecase

import (
    "practice02/internal/domain"
    "practice02/internal/domain/repository"
)

type TodoUsecase interface {
    Create(todo *domain.Todo) error
    GetAll() ([]domain.Todo, error)
    GetByID(id uint) (*domain.Todo, error)
    Update(todo *domain.Todo) error
    Delete(id uint) error
}

type todoUsecase struct {
    todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
    return &todoUsecase{todoRepo}
}

func (u *todoUsecase) Create(todo *domain.Todo) error {
    return u.todoRepo.Create(todo)
}

func (u *todoUsecase) GetAll() ([]domain.Todo, error) {
    return u.todoRepo.GetAll()
}

func (u *todoUsecase) GetByID(id uint) (*domain.Todo, error) {
    return u.todoRepo.GetByID(id)
}

func (u *todoUsecase) Update(todo *domain.Todo) error {
    return u.todoRepo.Update(todo)
}

func (u *todoUsecase) Delete(id uint) error {
    return u.todoRepo.Delete(id)
}