package users

import (
	"book_crud_ca/bussinesses"
	"book_crud_ca/helper/encrypt"
	"context"
	"strings"
	"time"
)

type userUseCase struct {
	Repository     Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repoUser Repository, timeout time.Duration) UseCase {
	return &userUseCase{
		Repository: repoUser,
	}
}

func (uc *userUseCase) Register(ctx context.Context, userDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedUser, err := uc.Repository.GetByUsername(ctx, userDomain.Username)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}

	if existedUser != (Domain{}) {
		return bussinesses.ErrDuplicateData
	}

	userDomain.Password, err = encrypt.Hash(userDomain.Password)
	if err != nil {
		return bussinesses.ErrInternalServer
	}

	if errInsert := uc.Repository.Insert(ctx, userDomain); errInsert != nil {
		return errInsert
	}

	return nil
}
