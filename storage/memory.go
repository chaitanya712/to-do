package storage

import (
	"errors"
	"sync"
	"to-do/models"
)

type MemoryStore struct {
	mu     sync.Mutex
	todos  map[int]models.Todo
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		todos:  make(map[int]models.Todo),
		nextID: 1,
	}
}

func (s *MemoryStore) Create(todo models.Todo) models.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo.ID = s.nextID
	s.todos[s.nextID] = todo
	s.nextID++

	return todo
}

func (s *MemoryStore) Get(id int) (models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, exists := s.todos[id]
	if !exists {
		return models.Todo{}, errors.New("task not found")
	}
	return todo, nil
}

func (s *MemoryStore) Update(id int, todo models.Todo) (models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.todos[id]
	if exists {
		return models.Todo{}, errors.New("task not found")
	}

	todo.ID = id
	s.todos[id] = todo

	return todo, nil
}

func (s *MemoryStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.todos[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(s.todos, id)
	return nil
}
