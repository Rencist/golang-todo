package repository

import (
	"Rencist/golang-todo/common"
	"Rencist/golang-todo/entity"
	"database/sql"
	"net/http"
)

type TodoRepository interface {
	GetAllTodo(w http.ResponseWriter, r *http.Request) ([]entity.Todo, error)
	CreateTodo(w http.ResponseWriter, r *http.Request, todo entity.Todo) (entity.Todo, error)
	GetTodoByID(w http.ResponseWriter, r *http.Request, todoID uint64) (entity.Todo, error)
	DeleteTodo(w http.ResponseWriter, r *http.Request, todoID uint64) (entity.Todo, error)
}

type todoConnection struct {
	connection *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoConnection{
		connection: db,
	}
}

func(db *todoConnection) CreateTodo(w http.ResponseWriter, r *http.Request, todo entity.Todo) (entity.Todo, error) {
	rows, err := db.connection.Query("INSERT INTO todo (todo) VALUES($1) RETURNING id, todo", todo.Todo)
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Menambahkan Todo", err.Error(), common.EmptyObj{})
	}
	rows.Next()

	result := entity.Todo{}

	rows.Scan(&result.ID, &result.Todo)
	return result, nil
}

func(db *todoConnection) GetAllTodo(w http.ResponseWriter, r *http.Request) ([]entity.Todo, error) {
	rows, err := db.connection.Query("SELECT * FROM TODO")
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Mengambil Data Todo", err.Error(), common.EmptyObj{})
	}
	defer rows.Close()
	result := entity.Todo{}
	arrresult := []entity.Todo{}
	for rows.Next() {
		rows.Scan(&result.ID, &result.Todo, &result.CreatedAt)
		if err != nil {
			common.BuildErrorResponse(w, "Gagal Mengambil Data Todo", err.Error(), common.EmptyObj{})
		}
		arrresult = append(arrresult, result)
	}
	err = rows.Err()
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Mengambil Data Todo", err.Error(), common.EmptyObj{})
	}
	return arrresult, nil
}

func(db *todoConnection) GetTodoByID(w http.ResponseWriter, r *http.Request, todoID uint64) (entity.Todo, error) {
	rows, err := db.connection.Query("SELECT * FROM TODO WHERE id = $1", todoID)
	rows.Next()
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Mengambil Data Todo", err.Error(), common.EmptyObj{})
	}
	defer rows.Close()
	result := entity.Todo{}
	rows.Scan(&result.ID, &result.Todo, &result.CreatedAt)
	err = rows.Err()
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Mengambil Data Todo", err.Error(), common.EmptyObj{})
	}
	return result, nil
}

func(db *todoConnection) DeleteTodo(w http.ResponseWriter, r *http.Request, todoID uint64) (entity.Todo, error) {
	rows, err := db.connection.Query("DELETE FROM TODO WHERE id = $1", todoID)
	rows.Next()
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Menghapus Todo", err.Error(), common.EmptyObj{})
	}
	defer rows.Close()
	result := entity.Todo{}
	rows.Scan(&result.ID, &result.Todo, &result.CreatedAt)
	err = rows.Err()
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Menghapus Todo", err.Error(), common.EmptyObj{})
	}
	return result, nil
}