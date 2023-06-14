package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-shop-app/configs"
)

func DBConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456/shop")
	if err != nil {
		WriteErrorLog(err)
	}
	configs.DBConfig(db)
	return db, nil
}
