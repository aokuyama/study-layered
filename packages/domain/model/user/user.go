package user

type User struct {
	id UserID
}
type UserWithPassword struct {
	User     User
	Password Password
}

func NewUser(id *UserID) (*User, error) {
	c := User{*id}
	return &c, nil
}

func (e *User) ID() *UserID {
	return &e.id
}

func (e *User) Identical(c *User) bool {
	return e.ID().Equals(c.ID().UUID)
}
