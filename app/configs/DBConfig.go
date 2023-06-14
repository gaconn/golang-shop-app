package configs

import (
	"database/sql"
	"time"
)

func DBConfig(db *sql.DB) *sql.DB {
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
