package request

import "book_crud_ca/bussinesses/users"

type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
	}
}
