package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/pedrohscramos/teste-frete-rapido/database/mysql"
	services "github.com/pedrohscramos/teste-frete-rapido/services/quote"
)

func QuotesRoutes(router *chi.Mux, db *mysql.DB) {
	// Quotes
	Service := services.NewGormRepository(db)
	quoteHandler := quote.NewQuoteHandler(Service)

	router.Post("/quotes", quoteHandler.InsertQuote)

}
