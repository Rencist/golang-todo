package main

import (
	"Rencist/golang-todo/config"
	"Rencist/golang-todo/controller"
	"Rencist/golang-todo/repository"
	"Rencist/golang-todo/routes"
	"Rencist/golang-todo/service"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	var(
		db *sql.DB = config.SetupDatabaseConnection()
		todoRepository repository.TodoRepository = repository.NewTodoRepository(db)
		todoService service.TodoService = service.NewTodoService(todoRepository)
		todoController controller.TodoController = controller.NewTodoController(todoService)
	)
	routes.TodoRoutes(todoController)
	fmt.Println("starting web server at http://localhost:8000/")
	http.ListenAndServe("127.0.0.1:8000", nil)
}