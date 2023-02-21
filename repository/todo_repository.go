package repository

import (
	"Rencist/golang-todo/entity"
	"database/sql"
	"log"
)

type TodoRepository interface {
	GetAllTodo() ([]entity.Todo, error)
	CreateTodo(todo entity.Todo) (entity.Todo, error)
	GetTodoByID(todoID uint64) (entity.Todo, error)
	DeleteTodo(todoID uint64) (entity.Todo, error)
}

type todoConnection struct {
	connection *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoConnection{
		connection: db,
	}
}

func(db *todoConnection) CreateTodo(todo entity.Todo) (entity.Todo, error) {
	rows, err := db.connection.Query("INSERT INTO todo (todo) VALUES($1) RETURNING id, todo", todo.Todo)
	if err != nil {
		return entity.Todo{}, err
	}
	rows.Next()

	result := entity.Todo{}

	rows.Scan(&result.ID, &result.Todo)
	return result, nil
}

func(db *todoConnection) GetAllTodo() ([]entity.Todo, error) {
	rows, err := db.connection.Query("SELECT * FROM TODO")
	if err != nil {
		return []entity.Todo{}, err
	}
	defer rows.Close()
	result := entity.Todo{}
	arrresult := []entity.Todo{}
	for rows.Next() {
		rows.Scan(&result.ID, &result.Todo, &result.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		arrresult = append(arrresult, result)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return arrresult, nil
}

func(db *todoConnection) GetTodoByID(todoID uint64) (entity.Todo, error) {
	rows, err := db.connection.Query("SELECT * FROM TODO WHERE id = $1", todoID)
	rows.Next()
	if err != nil {
		return entity.Todo{}, err
	}
	defer rows.Close()
	result := entity.Todo{}
	rows.Scan(&result.ID, &result.Todo, &result.CreatedAt)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func(db *todoConnection) DeleteTodo(todoID uint64) (entity.Todo, error) {
	rows, err := db.connection.Query("DELETE FROM TODO WHERE id = $1", todoID)
	rows.Next()
	if err != nil {
		return entity.Todo{}, err
	}
	defer rows.Close()
	result := entity.Todo{}
	rows.Scan(&result.ID, &result.Todo, &result.CreatedAt)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}