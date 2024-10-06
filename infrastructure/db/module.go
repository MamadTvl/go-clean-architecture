package db

import (
	"go.uber.org/fx"
)

var DbModule = fx.Module("DBModule",
	fx.Options(fx.Provide(NewDatabase), fx.Invoke(Migrate)),
)
