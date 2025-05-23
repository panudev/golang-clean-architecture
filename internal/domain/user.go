package domain

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserRepository interface {
	GetByID(id uint) (*User, error)
}
