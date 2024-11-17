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

// Project Repository

func (s *Storage) CreateProject(p *Project) (*Project, error) {
	result, err := s.db.Exec("INSERT into Project(name) values (?)", p.name)

	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	p.ID = id
	return p, nil

}

func (s *Storage) GetProject(id string) (*Project, error) {
	var p Project
	err := s.db.QueryRow("SELECT id, name, createdAt FROM projects WHERE id = ?", id).
		Scan(&p.ID, &p.name, &p.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("project with id %s not found", id)
		}
		return nil, err
	}

	return &p, nil

}

func (s *Storage) DeleteProject(id string) error {
	_, err := s.db.Exec("DELETE FROM projects WHERE id = ? ", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("project with id %s not found", id)
		}
		return err
	}

	return nil
}

// Task Repository

func (s *Storage) CreateTask(t *Task) (*Task, error) {
	result, err := s.db.Exec("INSERT into Tasks (name, status, project_id, assigned_to) VALUES(?,?,?,?,?)",
		t.Name, t.Status, t.ProjectID, t.AssignedToID
		)

	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	t.ID = id
	return t, nil

}

func (s *Storage) GetTask(id string) (*Task, error) {
	var t Task
	err := s.db.Exec("SELECT id, name, status, project_id, assigned_to, createdAt FROM tasks WHERE id = ?", id).
	.Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedToID, &t.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task with id %s not found", id)
		}
		return nil, err
	}

	return &t, nil

}
