package service

import (
	"tieu-minh/pkg/router"
	"tieu-minh/service/index"
	"tieu-minh/service/review"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {
	// Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)
	router.Router.Mount(router.RouterBasePath+"/todos", review.TodoServiceSubRouter)
}
