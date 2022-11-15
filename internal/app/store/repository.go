package store

import "github.com/BTkachenko/gotest/internal/app/model"


//UserRepository
type UserRepository interface {
	Create(* model.User) error
	FindByEmail(string) (*model.User, error)
}