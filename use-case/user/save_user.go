package user_interactor

import (
	"clean-architecture/domain/model"
	"clean-architecture/infrastructure/repository"
	bcrypt "clean-architecture/infrastructure/service/crypto"
	metrics "clean-architecture/infrastructure/service/prometheus"
)

type SaveUserUseCase struct {
	repo    repository.UserRepository
	bcrypt  bcrypt.Adapter
	metrics metrics.Metrics
}

func NewSaveUserUseCase(repo repository.UserRepository, bcrypt bcrypt.Adapter, metrics metrics.Metrics) *SaveUserUseCase {
	return &SaveUserUseCase{repo, bcrypt, metrics}
}

func (u *SaveUserUseCase) SaveUser(username, password string) (*model.User, error) {
	hashedPassword, err := u.bcrypt.GenerateHash(password)
	if err != nil {
		return nil, err
	}
	user, err := u.repo.Create(username, hashedPassword)
	if err != nil {
		return nil, err
	}
	u.metrics.IncrementUserCreation()
	return user, nil
}
