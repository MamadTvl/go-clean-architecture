package metrics

import "go.uber.org/fx"

var PrometheusMetricsModule = fx.Options(
	fx.Provide(NewPrometheusMetrics),
)
