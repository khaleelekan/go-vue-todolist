package repository

import "todo-app/internal/domain"

type TodoRepository interface {
	GetAll(userID int) ([]*domain.Todo, error)
	GetByID(id, userID int) (*domain.Todo, error)
	Create(todo *domain.Todo) (*domain.Todo, error)
	Update(todo *domain.Todo) (*domain.Todo, error)
	Delete(id, userID int) error
}
