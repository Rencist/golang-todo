package routes

import (
	"Rencist/golang-todo/controller"
	"net/http"
)

func TodoRoutes(TodoController controller.TodoController) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		TodoController.CreateTodo(w, r)
	})
}