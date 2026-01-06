package storage

import (
	"database/sql"
	"fmt"

	"github.com/Aytaditya/splitnest-expense-service/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	DB *sql.DB
}

func ConnectDB(config *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", config.Storagepath)
	if err != nil {
		return nil, err
	}
	_, err2 := db.Exec(`CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
  		group_id INTEGER NOT NULL,
  		paid_by INTEGER NOT NULL,
  		amount INTEGER NOT NULL,
  		created_at DATETIME DEFAULT CURRENT_TIMESTAMP)`)
	if err2 != nil {
		return nil, err2
	}

	_, err3 := db.Exec(`CREATE TABLE IF NOT EXISTS expense_splits(
		expense_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		amount INTEGER NOT NULL)`)
	if err3 != nil {
		return nil, err3
	}

	_, err4 := db.Exec(`CREATE TABLE IF NOT EXISTS balances (
    group_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    PRIMARY KEY (group_id, user_id)
	)`)
	if err4 != nil {
		return nil, err4
	}
	fmt.Println("Connected to SQLite database successfully")
	return &Sqlite{DB: db}, nil
}

func (s *Sqlite) ManageExpenses() error {
	return nil
}
