package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	
	const schema = `

	PRAGMA foreign_keys = ON;
	
	CREATE TABLE IF NOT EXISTS cars (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		brand TEXT NOT NULL,
		model TEXT NOT NULL,
		year INTEGER,
		rent INTEGER,
		photo TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE,
		birthday TIMESTAMP NOT NULL,
		is_admin BOOL,
		license_id INTEGER,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (license_id) REFERENCES licenses(id) ON DELETE SET NULL
	);

	CREATE TABLE IF NOT EXISTS licenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		number TEXT UNIQUE,
		issuance_date TIMESTAMP NOT NULL,
		expiration_date TIMESTAMP NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS bookings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		id_car INTEGER NOT NULL,
		id_user INTEGER NOT NULL,
		start_day TIMESTAMP NOT NULL,    
		end_day TIMESTAMP NOT NULL,
		daily_cost INTEGER NOT NULL,
		status TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	`

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}
	
	return db, db.Ping()
}