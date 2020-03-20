package user

// Repository is
type Repository interface {
	FindByID(id ID) (*User, error)
}
