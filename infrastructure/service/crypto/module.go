package bcrypt

import "go.uber.org/fx"

var BcryptModule = fx.Module("BcryptModule", fx.Options(fx.Provide(NewBcrypt)))
