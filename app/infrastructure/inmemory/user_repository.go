package inmemory

import user "worksample-go-layered-architecture/app/domain/user_aggregate"

// UserRepository is
type UserRepository struct {
}

// NewUserRepository is
func NewUserRepository() *UserRepository {
	return new(UserRepository)
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
