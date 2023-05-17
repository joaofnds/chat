package user

import "context"

type User struct {
	ID   string `json:"id"  gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name string `json:"name"`
}

type Probe interface {
	FailedToCreateUser(error)
	FailedToDeleteAll(error)
	FailedToFindByName(error)
	FailedToRemoveUser(error, User)
	UserCreated()
}

type Repository interface {
	CreateUser(context.Context, *User) error
	All(context.Context) ([]User, error)
	FindByName(context.Context, string) (User, error)
	Delete(context.Context, User) error
	DeleteAll(context.Context) error
}
