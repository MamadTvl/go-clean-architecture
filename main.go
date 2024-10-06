package main

import (
	"clean-architecture/infrastructure/config"
	user_controller "clean-architecture/infrastructure/controller/user"
	"clean-architecture/infrastructure/http"
	"clean-architecture/infrastructure/logger"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var AppModule = fx.Module("AppModule",
	config.ConfigModule,
	logger.LoggerModule,
	http.HttpModule,
	user_controller.UserControllerModule,
)

func main() {
	fx.New(AppModule, fx.WithLogger(func(logger logger.Logger) fxevent.Logger {
		return logger.GetFxLogger()
	})).Run()
}
