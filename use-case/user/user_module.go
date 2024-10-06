package user_interactor

import (
	"clean-architecture/infrastructure/repository"
	bcrypt "clean-architecture/infrastructure/service/crypto"

	"go.uber.org/fx"
)

var UserModule = fx.Module("UserModule", bcrypt.BcryptModule, repository.RepositoriesModule, fx.Options(fx.Provide(NewSaveUserUseCase)))
