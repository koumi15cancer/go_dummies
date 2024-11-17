package main

import (
	"errors"
)

var errEmailRequired = errors.New("Email is required")
var errFirstNameRequired = errors.New("First Name is required")
var errLastNameRequired = errors.New("Last Name is required")
var errPasswordRequired = errors.New("Password is required")

type UserService struct {
	store Store
}

func NewUserService(s Store) *UserService {
	return &UserService{store: s}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/user/register", s.handleUserRegister).Methods("POST")
	r.HandleFunc("/user/login", s.handleUserLogin).Methods("POST")
}
