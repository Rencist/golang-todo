package service

import (
	"Rencist/golang-todo/entity"
	"Rencist/golang-todo/repository"
)

type TodoService interface {
	GetAllTodo() ([]entity.Todo, error)
	CreateTodo(todo entity.Todo) (entity.Todo, error) 
	GetTodoByID(todoID uint64) (entity.Todo, error)
	// DeleteTodo(ctx context.Context, todoID uint64) (entity.Todo, error)
}

func NewTodoService(tr repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: tr,
	}
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func(ts *todoService) CreateTodo(todo entity.Todo) (entity.Todo, error) {
	return ts.todoRepository.CreateTodo(todo)
}

func(ts *todoService) GetAllTodo() ([]entity.Todo, error) {
	return ts.todoRepository.GetAllTodo()
}

func(ts *todoService) GetTodoByID(todoID uint64) (entity.Todo, error) {
	return ts.todoRepository.GetTodoByID(todoID)
}