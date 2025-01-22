package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(dbDriver string, dbName string) {
	var err error
	DB, err = sql.Open(dbDriver, dbName)

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id BIGSERIAL PRIMARY KEY,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic(err)
	}

	createEventTable := `
		CREATE TABLE IF NOT EXISTS events (
			id BIGSERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime TIMESTAMP NOT NULL,
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id)
			)
			`

	_, err = DB.Exec(createEventTable)

	if err != nil {
		panic(err)
	}

	createRegistrationTable := `
			CREATE TABLE IF NOT EXISTS registrations (
				id BIGSERIAL PRIMARY KEY,
				event_id INTEGER,
				user_id INTEGER,
				FOREIGN KEY (event_id) REFERENCES events(id),
				FOREIGN KEY (user_id) REFERENCES users(id)
		
		)
	`

	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		panic(err)
	}

}
