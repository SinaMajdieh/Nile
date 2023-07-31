// Package database is an implementation of sqlite3 database and consists of function
// to query on the specific database designed for this project
package database

// Sqlite database initialization

import (
	"Nile/internal/config"
	"Nile/internal/models"
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var logger = log.New(os.Stdout, "SQLITE: ", log.Ldate|log.Ltime)

var Db models.DataBase

type SqliteDb struct {
	db *sql.DB
}

// Initialize Sqlite database and opening it
func InitSqliteDb(config *config.Config) *SqliteDb {
	sqliteDb := SqliteDb{
		db: nil,
	}

	// Opening the database
	// Catching the error
	err := sqliteDb.Open(config.SqlitePath)
	if nil != err {
		logger.Fatal("Could not open sqlite3 database", err)
		return nil
	}

	// There was no errors
	logger.Println("Successfully opened " + config.SqlitePath)
	return &sqliteDb
}

// Opening the sqlite database
func (db *SqliteDb) Open(path string) error {
	dbFile, err := sql.Open("sqlite3", path)
	db.db = dbFile
	return err
}
