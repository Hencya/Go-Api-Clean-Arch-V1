package driver

import (
	"book_crud_ca/bussiness/users"
	usersDB "book_crud_ca/driver/databases/users"
	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) users.Repository {
	return usersDB.NewMySqlRepository(conn)
}
