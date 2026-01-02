package storage

import (
	"database/sql"
	"fmt"

	"github.com/Aytaditya/splitnest-user-service/internal/config"
	"github.com/Aytaditya/splitnest-user-service/internal/middleware"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type Sqlite struct {
	DB *sql.DB
}

func ConnectDB(config *config.Config) (*Sqlite, error) {
	fmt.Println("Connecting to database at path:", config.StoragePath)
	db, err := sql.Open("sqlite3", config.StoragePath)
	if err != nil {
		return nil, err
	}
	_, err1 := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL)`)

	if err1 != nil {
		return nil, err1
	}
	return &Sqlite{DB: db}, nil
}

func (s *Sqlite) RegisterUser(username, email, password string) (int64, string, error) {

	if username == "" || email == "" || password == "" {
		return 0, "", fmt.Errorf("username, email, and password cannot be nil")
	}
	fmt.Println("Registering user:", username, email, password)
	hashedPassword, err1 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err1 != nil {
		return 0, "", err1
	}
	stmt, err := s.DB.Prepare(`INSERT INTO users (username,email,password) VALUES (?,?,?)`)
	if err != nil {
		return 0, "", err
	}
	defer stmt.Close()
	res, err2 := stmt.Exec(username, email, string(hashedPassword))
	if err2 != nil {
		return 0, "", err2
	}
	id, err3 := res.LastInsertId()
	if err3 != nil {
		return 0, "", err3
	}
	token, err4 := middleware.CreateToken(id, email)
	if err4 != nil {
		return 0, "", err4
	}
	return id, token, nil
}
