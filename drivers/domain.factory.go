package drivers

import (
	"book_crud_ca/bussinesses/users"
	usersDB "book_crud_ca/drivers/databases/users"
	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) users.Repository {
	return usersDB.NewMySqlRepository(conn)
}
