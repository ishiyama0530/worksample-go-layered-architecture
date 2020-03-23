package user

// CreateUsecase is
type CreateUsecase interface {
	Handle(inp CreateInputData) (*CreateOutputData, error)
}

// CreateInputData is
type CreateInputData struct {
	Name string
}

// CreateOutputData is
type CreateOutputData struct {
	ID   string
	Name string
}

// NewCreateInputData is
func NewCreateInputData(name string) *CreateInputData {
	inp := new(CreateInputData)
	inp.Name = name
	return inp
}

// NewCreateOutputData is
func NewCreateOutputData(id, name string) *CreateOutputData {
	out := new(CreateOutputData)
	out.ID = id
	out.Name = name
	return out
}
