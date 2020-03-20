package user

import (
	"worksample-go-layered-architecture/app/domain/user"
	uc "worksample-go-layered-architecture/app/usecase/user"
)

// GetInteractor is
type GetInteractor struct {
	repo user.Repository
}

// NewGetInteractor is
func NewGetInteractor(repo user.Repository) *GetInteractor {
	interacotr := new(GetInteractor)
	interacotr.repo = repo
	return interacotr
}

// Handle is
func (interacotr *GetInteractor) Handle(inp uc.GetInputData) (*uc.GetOutputData, error) {
	id, err := user.NewID(inp.ID)
	if err != nil {
		return nil, err
	}
	user, err := interacotr.repo.FindByID(*id)
	if err != nil {
		return nil, err
	}

	out := new(uc.GetOutputData)
	out.ID = user.ID.Value()
	out.Name = user.Name.Value()
	return out, nil
}
