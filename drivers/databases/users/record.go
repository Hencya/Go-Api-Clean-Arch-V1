package users

import (
	"book_crud_ca/bussiness/users"
	"time"
)

type Users struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (rec *Users) ToDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		Username:  rec.Username,
		Password:  rec.Password,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func FromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:        userDomain.ID,
		Username:  userDomain.Username,
		Password:  userDomain.Password,
		Name:      userDomain.Name,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
