package repository

import (
	"Rencist/golang-todo/common"
	"Rencist/golang-todo/entity"
	"database/sql"
	"net/http"
)

type TodoRepository interface {
	CreateTodo(w http.ResponseWriter, r *http.Request, todo entity.Todo) (entity.Todo, error)
	GetAllTodo(w http.ResponseWriter, r *http.Request) ([]entity.Todo, error)
	MarkDone(w http.ResponseWriter, r *http.Request, todoID uint64) (entity.Todo, error)
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
	todo.IsDone = false
	rows, err := db.connection.Query("INSERT INTO todo (todo, is_done) VALUES($1, $2) RETURNING id, todo, is_done, created_at", todo.Todo, todo.IsDone)
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Menambahkan Todo", err.Error(), common.EmptyObj{})
	}
	rows.Next()

	result := entity.Todo{}

	rows.Scan(&result.ID, &result.Todo, &result.IsDone, &result.CreatedAt)
	return result, nil
}

func(db *todoConnection) GetAllTodo(w http.ResponseWriter, r *http.Request) ([]entity.Todo, error) {
	rows, err := db.connection.Query("SELECT id, todo, is_done, created_at FROM TODO")
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Mengambil Data Todo", err.Error(), common.EmptyObj{})
	}
	defer rows.Close()
	result := entity.Todo{}
	arrresult := []entity.Todo{}
	for rows.Next() {
		rows.Scan(&result.ID, &result.Todo, &result.IsDone, &result.CreatedAt)
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

func(db *todoConnection) MarkDone(w http.ResponseWriter, r *http.Request, todoID uint64) (entity.Todo, error) {
	rows, err := db.connection.Query("UPDATE todo SET is_done = true WHERE id = $1 RETURNING id, todo, is_done, created_at", todoID)
	// fmt.Println(rows.Next())
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Menandai Selesai Todo", err.Error(), common.EmptyObj{})
	}
	rows.Next()
	defer rows.Close()
	result := entity.Todo{}
	rows.Scan(&result.ID, &result.Todo, &result.IsDone, &result.CreatedAt)
	return result, nil
}

func(db *todoConnection) GetTodoByID(w http.ResponseWriter, r *http.Request, todoID uint64) (entity.Todo, error) {
	rows, err := db.connection.Query("SELECT id, todo, is_done, created_at FROM TODO WHERE id = $1", todoID)
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Mengambil Data Todo", err.Error(), common.EmptyObj{})
	}
	rows.Next()
	defer rows.Close()
	result := entity.Todo{}
	rows.Scan(&result.ID, &result.Todo, &result.IsDone, &result.CreatedAt)
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
	rows.Scan(&result.ID, &result.Todo, &result.IsDone, &result.CreatedAt)
	err = rows.Err()
	if err != nil {
		common.BuildErrorResponse(w, "Gagal Menghapus Todo", err.Error(), common.EmptyObj{})
	}
	return result, nil
}