package user_interactor

import (
	"clean-architecture/domain/model"
	"clean-architecture/infrastructure/repository"
	bcrypt "clean-architecture/infrastructure/service/crypto"
)

type SaveUserUseCase struct {
	repo   *repository.UserRepository
	bcrypt *bcrypt.Bcrypt
}

func NewSaveUserUseCase(repo *repository.UserRepository, bcrypt *bcrypt.Bcrypt) *SaveUserUseCase {
	return &SaveUserUseCase{repo, bcrypt}
}

func (uc *SaveUserUseCase) SaveUser(username string, password string) (*model.User, error) {
	hashedPassword, err := uc.bcrypt.GenerateHash(password)
	if err != nil {
		return nil, err
	}
	user, err := uc.repo.Create(username, hashedPassword)
	if err != nil {
		return nil, err
	}
	return user, nil
}
