package main

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

type Store interface {
	// User
	CreateUser(u *User) (*User, error)
	GetUserById(id string) (*User, error)

	// Projects
	CreateProject(p *Project) error
	GetProject(id string) (*Project, error)
	DeleteProject(id string) error

	// Tasks
	CreateTask(id string) (*Task, error)
	GetTask(id string) (*Task, error)
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

// User Repository
func (s *Storage) CreateUser(u *User) (*User, error) {
	result, err := s.db.Exec("INSERT INTO users (email, firstName, lastName, password) VALUES (?, ?, ?, ?)",
		u.Email, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	u.ID = id
	return u, nil
}

func (s *Storage) GetUserById(id string) (*User, error) {
	var u User
	err := s.db.QueryRow("SELECT id, email, firstName, lastName, createdAt FROM users WHERE id = ?", id).
		Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %s not found", id)
		}
		return nil, err
	}

	return &u, nil
}
