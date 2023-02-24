package routes

import (
	"Rencist/golang-todo/common"
	"Rencist/golang-todo/controller"
	"net/http"
)

func TodoRoutes(TodoController controller.TodoController) {
	http.HandleFunc("/todo/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			TodoController.CreateTodo(w, r)
		} else {
			common.BuildErrorResponse(w, "Endpoint Doesn`t Exist", "", common.EmptyObj{})
		}
	})

	http.HandleFunc("/todo/done", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			TodoController.MarkDone(w, r)
		} else {
			common.BuildErrorResponse(w, "Endpoint Doesn`t Exist", "", common.EmptyObj{})
		}
	})

	http.HandleFunc("/todo/index", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			TodoController.GetAllTodo(w, r)
		} else {
			common.BuildErrorResponse(w, "Endpoint Doesn`t Exist", "", common.EmptyObj{})
		}
	})
	
	http.HandleFunc("/todo/show", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			TodoController.GetTodoByID(w, r)
		} else {
			common.BuildErrorResponse(w, "Endpoint Doesn`t Exist", "", common.EmptyObj{})
		}
	})

	http.HandleFunc("/todo/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			TodoController.DeleteTodo(w, r)
		} else {
			common.BuildErrorResponse(w, "Endpoint Doesn`t Exist", "", common.EmptyObj{})
		}
	})
}