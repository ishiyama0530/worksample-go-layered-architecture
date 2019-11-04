package inmemory

import (
	"github.com/ishiyama0530/worksample-go-layered-architecture/app/domain/user"
)

// UserStore is
type UserStore struct {
}

// FindByID is
func (repo *UserStore) FindByID(id int) (user *user.User, err error) {
	user.ID = 1
	user.Name = "test"
	return
}
