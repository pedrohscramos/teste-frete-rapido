package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
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

	router.Get("/csrf", func(w http.ResponseWriter, r *http.Request) {
		if config.GetAPIConfig().Environment == "production" {
			session := r.Context().Value("session").(*sessions.Session)
			if session.IsNew {
				session.Options = &sessions.Options{
					Path:   "/",
					MaxAge: 86400,
				}
				session.Save(r, w)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"token": csrf.Token(r)})
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": "sandbox"})
	})

	QuotesRoutes(router, db)

}
