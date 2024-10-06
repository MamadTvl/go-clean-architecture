package user_controller

import (
	user_interactor "clean-architecture/use-case/user"

	"go.uber.org/fx"
)

var UserControllerModule = fx.Module("UserControllerModule",
	user_interactor.UserModule,
	fx.Options(
		fx.Provide(
			NewController,
			NewRoute,
		),
		fx.Invoke(RegisterRoute),
	),
)
