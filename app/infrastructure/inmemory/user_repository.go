package inmemory

import "worksample-go-layered-architecture/app/domain/user"

// UserRepository is
type UserRepository struct {
}

// FindByID is
func (*UserRepository) FindByID(id user.ID) (*user.User, error) {
	localID, err := user.NewID("id")
	if err != nil {
		return nil, err
	}
	name, err := user.NewName("name")
	if err != nil {
		return nil, err
	}
	return user.NewUser(*localID, *name), nil
}
