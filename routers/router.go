package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pedrohscramos/teste-frete-rapido/database/mysql"
	"github.com/pedrohscramos/teste-frete-rapido/middlewares"
)

func SetupRoute(db *mysql.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middlewares.Cors())
	// router.Use(middlewares.Csrf())

	RegisterRoutes(router, db)

	return router
}
