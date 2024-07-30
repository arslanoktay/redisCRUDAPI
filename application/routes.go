package application

import (
	"net/http"
	"redisCRUDAPI/handler"
	"redisCRUDAPI/repository/order"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() { // mux is handler multiplexer(çoğaltıcı)
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", a.loadOrderRoutes)

	a.router = router
}

func (a *App) loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{
		Repo: &order.RedisRepo{
			Client: a.rdb,
		},
	}

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetById)
	router.Put("/{id}", orderHandler.UpdateByID)
	router.Delete("/{id}", orderHandler.DeleteByID)
}
