package repository

import (
	"clean-architecture/infrastructure/db"

	"go.uber.org/fx"
)

var RepositoriesModule = fx.Module("RepositoriesModule", db.DbModule, fx.Options(fx.Provide(NewUserRepository)))
