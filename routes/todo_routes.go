package routes

import (
	"Rencist/golang-todo/controller"
	"net/http"
)

func TodoRoutes(TodoController controller.TodoController) {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			TodoController.CreateTodo(w, r)
		}
	})

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			TodoController.GetAllTodo(w, r)
		}
	})
	
	http.HandleFunc("/show", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			TodoController.GetTodoByID(w, r)
		}
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			TodoController.DeleteTodo(w, r)
		}
	})
}