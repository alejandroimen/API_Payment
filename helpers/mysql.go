package helpers

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewMySQLConnection() (*sql.DB, error) {
	var err error
	once.Do(func() {
		mysqlUser := os.Getenv("MYSQL_USER")
		mysqlPassword := os.Getenv("MYSQL_PASSWORD")
		mysqlHost := os.Getenv("MYSQL_HOST")
		mysqlPort := os.Getenv("MYSQL_PORT")
		mysqlDatabase := os.Getenv("MYSQL_DATABASE")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)

		db, err = sql.Open("mysql", dsn)
		if err != nil {
			err = fmt.Errorf("error conectando con MySQL: %w", err)
			return
		}

		if err = db.Ping(); err != nil {
			err = fmt.Errorf("error haciendo ping a MySQL: %w", err)
			return
		}
	})
	return db, err
}