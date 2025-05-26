package domain

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	GetByID(id uint) (*User, error)
}
