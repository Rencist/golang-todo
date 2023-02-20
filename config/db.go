package config

import "database/sql"

func SetupDatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgres://postgres:rencist@localhost:5432/todo")
	if err != nil {
		return nil, err
	}

	return db, nil
}
