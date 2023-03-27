package review

import (
	"tieu-minh/service/review/controller"

	"github.com/go-chi/chi"
)

var TodoServiceSubRouter = chi.NewRouter()

// Init package with sub-route for price service
func init() {
	TodoServiceSubRouter.Group(func(r chi.Router) {
		TodoServiceSubRouter.Get("/list-todo", controller.GetListTodo)
		TodoServiceSubRouter.Post("/add-todo", controller.AddTodo)
		TodoServiceSubRouter.Delete("/delete-todo", controller.DeleteTodo)
	})
}
