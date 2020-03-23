package user

import (
	"worksample-go-layered-architecture/app/domain/user"
	uc "worksample-go-layered-architecture/app/usecase/user"

	"github.com/google/uuid"
)

// CreateInteractor is
type CreateInteractor struct {
	repo user.Repository
}

// NewCreateInteractor is
func NewCreateInteractor(repo user.Repository) *CreateInteractor {
	interacotr := new(CreateInteractor)
	interacotr.repo = repo
	return interacotr
}

// Handle is
func (i *CreateInteractor) Handle(inp uc.CreateInputData) (*uc.CreateOutputData, error) {
	// required Transaction

	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	id, err := user.NewID(uuid.String())
	if err != nil {
		return nil, err
	}
	name, err := user.NewName(inp.Name)
	if err != nil {
		return nil, err
	}

	newUser, err := user.NewUser(*id, *name)
	if err != nil {
		return nil, err
	}

	err = i.repo.Save(*newUser)
	if err != nil {
		return nil, err
	}

	out := new(uc.CreateOutputData)
	out.ID = newUser.ID.Value()
	out.Name = newUser.Name.Value()
	return out, nil
}
