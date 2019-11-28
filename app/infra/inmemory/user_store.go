package inmemory

import (
	domain "github.com/ishiyama0530/worksample-go-layered-architecture/app/domain/user"
)

// UserStore is
type UserStore struct {
}

// FindByID is
func (repo *UserStore) FindByID(id int) (user *domain.User, err error) {
	user = &domain.User{
		ID:   id,
		Name: "ishiyama",
	}
	return
}
