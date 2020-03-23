package user

import (
	uc "worksample-go-layered-architecture/app/usecase/user"
)

// GetInteractor is
type GetInteractor struct {
}

// NewGetInteractor is
func NewGetInteractor() *GetInteractor {
	return new(GetInteractor)
}

// Handle is
func (i *GetInteractor) Handle(inp uc.GetInputData) (out *uc.GetOutputData, err error) {
	out = new(uc.GetOutputData)
	out.ID = "ID"
	out.Name = "NAME"
	return
}
