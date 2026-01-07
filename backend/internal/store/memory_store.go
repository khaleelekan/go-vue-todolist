package store

import (
	"errors"
	"sync"
	"time"
	"todo-app/internal/models"
)

var (
	ErrNotFound = errors.New("todo not found")
)

type TodoStore struct {
	mu     sync.RWMutex
	todos  map[int]*models.Todo
	nextID int
}

func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos:  make(map[int]*models.Todo),
		nextID: 1,
	}
}

func (s *TodoStore) GetAll() []*models.Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*models.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}
	return todos
}

func (s *TodoStore) GetByID(id int) (*models.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, exists := s.todos[id]
	if !exists {
		return nil, ErrNotFound
	}
	return todo, nil
}

func (s *TodoStore) Create(title string) *models.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := &models.Todo{
		ID:        s.nextID,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	s.todos[todo.ID] = todo
	s.nextID++
	return todo
}

func (s *TodoStore) Update(id int, title string, completed bool) (*models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, exists := s.todos[id]
	if !exists {
		return nil, ErrNotFound
	}

	if title != "" {
		todo.Title = title
	}
	todo.Completed = completed
	return todo, nil
}

func (s *TodoStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.todos[id]; !exists {
		return ErrNotFound
	}
	delete(s.todos, id)
	return nil
}
