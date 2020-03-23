package user

// Repository is
type Repository interface {
	Save(user User) error
}
