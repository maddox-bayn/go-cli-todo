package todo

import (
	"errors"
)

type TodoService struct {
	storage TodoStorage
}

func NewTodoService(storage TodoStorage) *TodoService {
	return &TodoService{storage: storage}
}

func (s *TodoService) AddTodo(title string) error {
	if title == "" {
		return errors.New("title can not be empty")
	}

	todos, err := s.storage.Load()
	if err != nil {
		return err
	}

	var idMax int
	for _, t := range todos {
		if idMax < t.ID {
			idMax = t.ID
		}
	}

	newTodo := Todo{
		ID:    idMax + 1,
		Title: title,
		Done:  false,
	}

	todos = append(todos, newTodo)
	return s.storage.Save(todos)
}
