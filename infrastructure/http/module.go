package http

import (
	"go.uber.org/fx"
)

var HttpModule = fx.Module("HttpModule", fx.Options(fx.Provide(NewRouter)))
