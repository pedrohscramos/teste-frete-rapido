package routers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pedrohscramos/teste-frete-rapido/config"
	"github.com/pedrohscramos/teste-frete-rapido/database/mysql"
	_ "github.com/pedrohscramos/teste-frete-rapido/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes(router *chi.Mux, db *mysql.DB) {
	//Configuração do Swagger
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Bem vindo, API Frete Rápido 👋!\nDocumentação: %s", config.GetAPIConfig().Documentation)))
	})

	QuotesRoutes(router, db)

}
