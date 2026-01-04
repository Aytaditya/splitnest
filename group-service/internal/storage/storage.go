package storage

import (
	"database/sql"
	"fmt"

	"github.com/Aytaditya/splitnest-group-service/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	DB *sql.DB
}

func ConnectDB(cfg *config.Config) (*Sqlite, error) {
	fmt.Println("Path for storage", cfg.StoragePath)
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	_, err1 := db.Exec(`CREATE TABLE IF NOT EXISTS groups (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_by INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP)`)
	if err1 != nil {
		return nil, err1
	}

	_, err2 := db.Exec(`CREATE TABLE IF NOT EXISTS group_members (
		group_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		joined_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (group_id, user_id)
		FOREIGN KEY (group_id) REFERENCES groups(id)
		)`)

	if err2 != nil {
		return nil, err2
	}

	return &Sqlite{DB: db}, nil
}

func (s *Sqlite) CreateGroup(name string, createdBy int64) (int64, error) {
	if name == "" || createdBy == 0 {
		return 0, fmt.Errorf("invalid group name or creator id")
	}
	stmt, err := s.DB.Prepare(`INSERT INTO groups (name,created_by) VALUES (?,?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(name, createdBy)
	if err2 != nil {
		return 0, err2
	}
	id, err3 := result.LastInsertId()
	if err3 != nil {
		return 0, err3
	}
	return id, nil
}

func (s *Sqlite) AddMembers(groupId, requesterId, userId int64) error {
	if groupId == 0 || userId == 0 || requesterId == 0 {
		return fmt.Errorf("invalid group id or user id")
	}
	// now we will validate that the group is created by requester
	stmt, err := s.DB.Prepare(`SELECT created_by FROM groups WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRow(groupId)
	var createdBy int64
	err3 := row.Scan(&createdBy)
	if err3 != nil {
		return err3
	}

	if createdBy != requesterId {
		return fmt.Errorf("only group owner can add members")
	}

	stmt2, err4 := s.DB.Prepare(`INSERT INTO group_members (group_id,user_id) VALUES (?,?)`)
	if err4 != nil {
		return err4
	}
	defer stmt2.Close()
	_, err5 := stmt2.Exec(groupId, userId)
	if err5 != nil {
		return err5
	}

	return nil
}
