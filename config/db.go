package config

import "database/sql"

func SetupDatabaseConnection() (*sql.DB) {
	db, err := sql.Open("pgx", "postgres://postgres:rencist@localhost:5432/todo")
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
