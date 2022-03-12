// Package models handles all database access/interaction
// During testing, we use the same connection
// All db access is done in a transaction via a closure
package models

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

// package level db connection
var db *gorm.DB

// InitDB connect and setup db
//TODO: refactor to use runtime environment mode (test, production, dev)
func InitDB() error {
	var err error
	//ignore if created and ping is ok
	if db != nil {
		var sqlDb *sql.DB
		if sqlDb, err = db.DB(); err == nil {
			if sqlDb.Ping() == nil {
				return nil
			}
		}
	}
	if err = godotenv.Load(".env"); err != nil {
		return err
	}
	dbUsername := os.Getenv("POSTGRES_USER")
	dbPwd := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_DB_HOST")
	dbPort := os.Getenv("POSTGRES_DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUsername, dbPwd, dbName, dbPort)
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	//	dbHost, dbUsername, dbPwd, dbName, dbPort)
	fmt.Println(dsn)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	var sqlDb *sql.DB
	sqlDb, err = db.DB()
	if err != nil {
		return err
	}

	// setup pool. ignore
	//sqlDb.SetMaxIdleConns(10)
	//sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	return sqlDb.Ping() //ensure db is reachable
}

func CloseDB() {
	sqlDb, _ := db.DB()
	sqlDb.Close()
}

// ExecDBFuncReadOnly run function in a readonly transaction
func ExecDBFuncReadOnly(fn func(tx *gorm.DB) error) error {
	return db.Begin(&sql.TxOptions{ReadOnly: true}).Transaction(fn)
}

// ExecDBFunc run function in a read/write transaction
func ExecDBFunc(fn func(tx *gorm.DB) error) error {
	return db.Transaction(fn)
}

func NewTx() *gorm.DB {
	return db.Begin()
}
