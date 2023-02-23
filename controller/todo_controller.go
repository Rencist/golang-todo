package controller

import (
	"Rencist/golang-todo/common"
	"Rencist/golang-todo/entity"
	"Rencist/golang-todo/service"
	"net/http"
	"strconv"
)

type TodoController interface {
	GetAllTodo(w http.ResponseWriter, r *http.Request)
	CreateTodo(w http.ResponseWriter, r *http.Request)
	GetTodoByID(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
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
	res, err := service.TodoService.CreateTodo(tc.todoService, w, r, todo)
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Menambahkan Todo", err.Error(), common.EmptyObj{})
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
	res, err := service.TodoService.GetAllTodo(tc.todoService, w, r)
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Mengambil Data Todo", err.Error(), common.EmptyObj{})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// temp, err := template.ParseFiles("views/index.html")
	// if err != nil {
	// 	common.BuildErrorResponse(w, "Gagal Mengambil Data Todo", err.Error(), common.EmptyObj{})
	// }
	// data := map[string]interface{}{
	// 	"data": res,
	// }
	// fmt.Fprintln(w, data)
	// temp.Execute(w, data)
	common.BuildResponse(w, true, "OK", res)
}

func(tc *todoController) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	todo_id := r.URL.Query().Get("id")
	lmao, _ := strconv.ParseUint(string(todo_id), 10, 64)
	res, err := service.TodoService.GetTodoByID(tc.todoService, w, r, lmao)
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Mengambil Data Todo", err.Error(), common.EmptyObj{})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if res.Todo == "" {
		common.BuildResponse(w, true, "OK", common.EmptyObj{})
	} else {
		// temp, err := template.ParseFiles("views/index.html")
		// if err != nil {
		// 	common.BuildErrorResponse(w, "Gagal Mengambil Data Todo", err.Error(), common.EmptyObj{})
		// }
		// data := map[string]interface{}{
		// 	"ID": res.ID,
		// 	"Todo": res.Todo,
		// 	"CreatedAt": res.CreatedAt,
		// }
		// temp.Execute(w, data)
		common.BuildResponse(w, true, "OK", res)
	}
}

func(tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todo_id := r.URL.Query().Get("id")
	lmao, _ := strconv.ParseUint(string(todo_id), 10, 64)
	_, err := service.TodoService.DeleteTodo(tc.todoService, w, r, lmao)
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Menghapus Todo", err.Error(), common.EmptyObj{})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	common.BuildResponse(w, true, "OK", common.EmptyObj{})
}
