package users

import (
	"book_crud_ca/bussiness/users"
	"context"
	"gorm.io/gorm"
)

type mysqlUsersRepository struct {
	Conn *gorm.DB
}

func NewMySqlRepository(conn *gorm.DB) users.Repository {
	return &mysqlUsersRepository{
		Conn: conn,
	}
}

func (nr *mysqlUsersRepository) Insert(ctx context.Context, userDomain *users.Domain) error {
	rec := FromDomain(*userDomain)

	result := nr.Conn.Create(&rec)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (nr *mysqlUsersRepository) GetByUsername(ctx context.Context, username string) (users.Domain, error) {
	rec := Users{}
	if err := nr.Conn.Where("username = ?", username).First(&rec).Error; err != nil {
		return users.Domain{}, err
	}
	return rec.ToDomain(), nil
}
