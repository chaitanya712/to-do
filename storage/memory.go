package storage

import (
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
