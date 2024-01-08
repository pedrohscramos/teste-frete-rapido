package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/pedrohscramos/teste-frete-rapido/database/mysql"
	"github.com/pedrohscramos/teste-frete-rapido/handlers/quotes"
	services "github.com/pedrohscramos/teste-frete-rapido/services/quotes"
)

func QuotesRoutes(router *chi.Mux, db *mysql.DB) {
	// Quotes
	Service := services.NewGormRepository(db)
	quoteHandler := quotes.NewQuoteHandler(Service)

	router.Post("/quote", quoteHandler.InsertQuote)
	router.Get("/metrics", quoteHandler.GetLastQuotes)

}
