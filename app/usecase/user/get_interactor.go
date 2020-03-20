package usecase

import (
	"worksample-go-layered-architecture/app/domain/user"
)

// GetUserInteractor is
type GetUserInteractor struct {
	repository user.Repository
}

// GetUserResponse is
type GetUserResponse struct {
	ID   int
	Name string
}

// NewGetUserInteractor is
func NewGetUserInteractor(repository user.Repository) GetUserInteractor {
	return GetUserInteractor{
		repository: repository,
	}
}

// GetUser is
func (interactor *GetUserInteractor) GetUser(id int) (res GetUserResponse, err error) {
	user, err := interactor.repository.FindByID(id)
	res.ID = user.ID
	res.Name = user.Name
	return
}
