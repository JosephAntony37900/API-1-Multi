package helpers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

// NewMySQLConnection devuelve una única instancia de la conexión a la base de datos.
func NewMySQLConnection() (*sql.DB, error) {
	var err error
	once.Do(func() {
		dsn := "admin:chocolate1804@tcp(database-1.c3g30jgzgku5.us-east-1.rds.amazonaws.com:3306)/multi_digisan"
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