package inmemory

import (
	"worksample-go-layered-architecture/app/domain/user"
)

// UserRepository is
type UserRepository struct {
}

// NewUserRepository is
func NewUserRepository() *UserRepository {
	return new(UserRepository)
}

// Save is
func (r *UserRepository) Save(user user.User) error {
	return nil
}
