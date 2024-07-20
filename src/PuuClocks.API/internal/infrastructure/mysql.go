package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL interface {
	Close() error
	Query(query string, args ...string) (*sql.Rows, error)
	QueryRow(query string, args ...string) *sql.Row
	Exec(query string, args ...string) (sql.Result, error)

	Health() error
}

type mySQL struct {
	DB *sql.DB
}

type MySQLConfig struct {
	DBName                   string
	Path                     string
	MaxIdleConns             int
	MaxOpenConns             int
	ConnMaxLifetimeInMinutes int
}

func NewMySQL(config MySQLConfig) (MySQL, error) {
	db, err := sql.Open(config.DBName, config.Path)
	if err != nil {
		return nil, fmt.Errorf("mysql - couldn't open db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("mysql - connection is not healthy: %w", err)
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(config.ConnMaxLifetimeInMinutes))
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)

	return mySQL{
		DB: db,
	}, nil
}

func (m mySQL) Close() error {
	return m.DB.Close()
}

func (m mySQL) Query(query string, args ...string) (*sql.Rows, error) {
	return m.DB.Query(query, args)
}

func (m mySQL) QueryRow(query string, args ...string) *sql.Row {
	return m.DB.QueryRow(query, args)
}

func (m mySQL) Exec(query string, args ...string) (sql.Result, error) {
	return m.DB.Exec(query, args)
}

func (m mySQL) Health() error {
	return m.DB.Ping()
}
