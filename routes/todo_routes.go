package routes

import (
	"Rencist/golang-todo/controller"
	"net/http"
)

func TodoRoutes(TodoController controller.TodoController) {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		TodoController.CreateTodo(w, r)
	})

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		TodoController.GetAllTodo(w, r)
	})
	
	http.HandleFunc("/show", func(w http.ResponseWriter, r *http.Request) {
		TodoController.GetTodoByID(w, r)
	})
}