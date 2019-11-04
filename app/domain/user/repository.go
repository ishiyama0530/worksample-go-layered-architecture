package user

// Repository is
type Repository interface {
	FindByID(id int) (*User, error)
}
