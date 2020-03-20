package user

// GetUsecase is
type GetUsecase interface {
	Handle(inp GetInputData) (*GetOutputData, error)
}

// GetInputData is
type GetInputData struct {
	ID string
}

// GetOutputData is
type GetOutputData struct {
	ID   string
	Name string
}

// NewGetInputData is
func NewGetInputData(id string) *GetInputData {
	inp := new(GetInputData)
	inp.ID = id
	return inp
}

// NewGetOutputData is
func NewGetOutputData(id, name string) *GetOutputData {
	out := new(GetOutputData)
	out.ID = id
	out.Name = name
	return out
}
