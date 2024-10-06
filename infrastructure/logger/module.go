package logger

import (
	"go.uber.org/fx"
)

var LoggerModule = fx.Module("LoggerModule", fx.Options(fx.Provide(GetLogger)))
