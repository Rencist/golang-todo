package service

import (
	"Rencist/golang-todo/entity"
	"Rencist/golang-todo/repository"
	"net/http"
)

type TodoService interface {
	GetAllTodo(w http.ResponseWriter, r *http.Request) ([]entity.Todo, error)
	CreateTodo(w http.ResponseWriter, r *http.Request, todo entity.Todo) (entity.Todo, error)
	MarkDone(w http.ResponseWriter, r *http.Request, todoID uint64) (entity.Todo, error)
	GetTodoByID(w http.ResponseWriter, r *http.Request, todoID uint64) (entity.Todo, error)
	DeleteTodo(w http.ResponseWriter, r *http.Request, todoID uint64) (entity.Todo, error)
}

func NewTodoService(tr repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: tr,
	}
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func(ts *todoService) CreateTodo(w http.ResponseWriter, r *http.Request, todo entity.Todo) (entity.Todo, error) {
	return ts.todoRepository.CreateTodo(w, r, todo)
}

func(ts *todoService) GetAllTodo(w http.ResponseWriter, r *http.Request) ([]entity.Todo, error) {
	return ts.todoRepository.GetAllTodo(w, r)
}

func(ts *todoService) GetTodoByID(w http.ResponseWriter, r *http.Request, todoID uint64) (entity.Todo, error) {
	return ts.todoRepository.GetTodoByID(w, r, todoID)
}

func(ts *todoService) MarkDone(w http.ResponseWriter, r *http.Request, todoID uint64) (entity.Todo, error) {
	return ts.todoRepository.MarkDone(w, r, todoID)
}

func(ts *todoService) DeleteTodo(w http.ResponseWriter, r *http.Request, todoID uint64) (entity.Todo, error) {
	return ts.todoRepository.DeleteTodo(w, r, todoID)
}