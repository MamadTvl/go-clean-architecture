package http

import (
	"clean-architecture/infrastructure/config"
	"clean-architecture/infrastructure/logger"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"net/http"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	lc fx.Lifecycle,
	logger logger.Logger,
	env *config.Config,
) Router {

	gin.SetMode(gin.DebugMode)

	httpRouter := gin.Default()

	gin.DefaultWriter = logger.GetGinLogger()

	httpRouter.GET("/metrics", gin.WrapH(promhttp.Handler()))

	httpRouter.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "clean architecture 📺 API Up and Running"})
	})
	lc.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		go func() {
			err := httpRouter.Run(":8000")
			if err != nil {
				logger.Error(err)
			}
		}()
		return nil
	}})

	return Router{
		httpRouter,
	}
}
