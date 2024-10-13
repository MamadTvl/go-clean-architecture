package user_controller

import "clean-architecture/infrastructure/http"

type Route struct {
	handler    http.Router
	controller *Controller
}

func NewRoute(
	handler http.Router,
	controller *Controller,
) *Route {
	return &Route{
		handler:    handler,
		controller: controller,
	}

}

// RegisterRoute Setup user routes
func RegisterRoute(r *Route) {
	api := r.handler.Group("/api")
	api.POST("/user", r.controller.CreateUser)
}
