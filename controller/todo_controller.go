package controller

import (
	"Rencist/golang-todo/common"
	"Rencist/golang-todo/entity"
	"Rencist/golang-todo/service"
	"fmt"
	"net/http"
	"strconv"
)

type TodoController interface {
	GetAllTodo(w http.ResponseWriter, r *http.Request)
	CreateTodo(w http.ResponseWriter, r *http.Request)
	GetTodoByID(w http.ResponseWriter, r *http.Request)
	// DeleteTodo(ctx *gin.Context)
}

type todoController struct {
	todoService service.TodoService
}

func NewTodoController(ts service.TodoService) TodoController {
	return &todoController{
		todoService: ts,
	}
}

func(tc *todoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	todo := entity.Todo{
		Todo: r.FormValue("todo"),
	}
	res, err := service.TodoService.CreateTodo(tc.todoService, todo)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := entity.Todo{
		ID: res.ID,
		Todo: res.Todo,
		CreatedAt: res.CreatedAt,
	}
	common.BuildResponse(w, true, "OK", data)
}

func(tc *todoController) GetAllTodo(w http.ResponseWriter, r *http.Request) {
	res, err := service.TodoService.GetAllTodo(tc.todoService)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	common.BuildResponse(w, true, "OK", res)
}

func(tc *todoController) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	todo_id := r.URL.Query().Get("id")
	lmao, _ := strconv.ParseUint(string(todo_id), 10, 64)
	res, err := service.TodoService.GetTodoByID(tc.todoService, lmao)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	common.BuildResponse(w, true, "OK", res)
}
