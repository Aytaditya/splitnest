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

// will write into expenses table and return expenseId
func (s *Sqlite) CreateExpense(groupId, userId, amount int64) (int64, error) {
	if groupId <= 0 || userId <= 0 || amount <= 0 {
		return 0, fmt.Errorf("invalid groupId, userId, or amount")
	}

	stmt, err := s.DB.Prepare(`INSERT INTO expenses (group_id, paid_by, amount) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	res, err2 := stmt.Exec(groupId, userId, amount)
	defer stmt.Close()
	if err2 != nil {
		return 0, err2
	}
	expenseId, err3 := res.LastInsertId()
	if err3 != nil {
		return 0, err3
	}
	return expenseId, nil
}

// will write into expense_splits table
func (s *Sqlite) CreateExpenseSplit(expenseId, memberId, splitAmount int64) {
	if expenseId <= 0 || memberId <= 0 || splitAmount <= 0 {
		return
	}
	stmt, err := s.DB.Prepare(`INSERT INTO expense_splits (expense_id, user_id, amount) VALUES (?, ?, ?)`)
	if err != nil {
		return
	}
	_, err2 := stmt.Exec(expenseId, memberId, splitAmount)
	defer stmt.Close()
	if err2 != nil {
		return
	}
	return
}

// will update balances table
func (s *Sqlite) UpdateBalance(groupId, userId, amount int64) {
	if groupId <= 0 || userId <= 0 {
		return
	}
	// check if entry exists
	var existingAmount int64
	err := s.DB.QueryRow(`SELECT amount FROM balances WHERE group_id = ? AND user_id = ?`, groupId, userId).Scan(&existingAmount)
	if err != nil {
		// insert new entry
		stmt, err2 := s.DB.Prepare(`INSERT INTO balances (group_id, user_id, amount) VALUES (?, ?, ?)`)
		if err2 != nil {
			return
		}
		_, err3 := stmt.Exec(groupId, userId, amount)
		defer stmt.Close()
		if err3 != nil {
			return
		}
	} else {
		// update existing entry
		newAmount := existingAmount + amount
		stmt, err2 := s.DB.Prepare(`UPDATE balances SET amount = ? WHERE group_id = ? AND user_id = ?`)
		if err2 != nil {
			return
		}
		_, err3 := stmt.Exec(newAmount, groupId, userId)
		defer stmt.Close()
		if err3 != nil {
			return
		}
	}
	return
}
