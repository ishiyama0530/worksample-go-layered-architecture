package user

// User is
type User struct {
	ID   ID
	Name Name
}

// NewUser is
func NewUser(id ID, name Name) *User {
	u := new(User)
	u.ID = id
	u.Name = name
	return u
}
