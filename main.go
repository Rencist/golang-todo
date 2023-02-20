package main

import (
	"Rencist/golang-todo/common"
	"Rencist/golang-todo/config"
	"Rencist/golang-todo/entity"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func CreateTodo(db *sql.DB, todo entity.Todo) (*entity.Todo, error) {
	rows, err := db.Query("INSERT INTO todo (todo) VALUES($1) RETURNING id, todo", todo.Todo)
	if err != nil {
		return nil, err
	}
	rows.Next()

	result := entity.Todo{}

	rows.Scan(&result.ID, &result.Todo)
	return &result, nil
}

func main() {
	db, err := config.SetupDatabaseConnection()
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	todo := entity.Todo{
		Todo: "lmao",
	}
	
	res, err := CreateTodo(db, todo)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// tell the client that the content type is json
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		data := entity.Todo{
			ID: res.ID,
			Todo: res.Todo,
			CreatedAt: res.CreatedAt,
		}
		common.BuildResponse(w, true, "OK", data)
	})
	
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe("127.0.0.1:8000", nil)
}