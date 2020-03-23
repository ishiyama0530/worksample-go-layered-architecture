package user

import (
	"github.com/pkg/errors"
)

// ID is
type ID struct {
	value string
}

// NewID is
func NewID(value string) (*ID, error) {
	if value == "" {
		return nil, errors.New("value must not be empty")
	}
	if len(value) > 20 {
		return nil, errors.New("value must not be exceed 20 characters")
	}

	id := new(ID)
	id.value = value

	return id, nil
}

// Value is
func (id *ID) Value() string {
	return id.value
}

// Equal is
func (id *ID) Equal(arg *ID) bool {
	return id.value == arg.Value()
}
