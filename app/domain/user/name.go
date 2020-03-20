package user

import (
	"github.com/pkg/errors"
)

// Name is
type Name struct {
	value string
}

// NewName is
func NewName(value string) (*Name, error) {
	if value == "" {
		return nil, errors.New("value must not be empty")
	}
	if len(value) > 12 {
		return nil, errors.New("value must not be exceed 12 characters")
	}

	name := new(Name)
	name.value = value

	return name, nil
}

// Value is
func (name *Name) Value() string {
	return name.value
}
