package main

import (
	"log"
	"net/http"

	"github.com/pedrohscramos/teste-frete-rapido/config"
	"github.com/pedrohscramos/teste-frete-rapido/database/mysql"
	"github.com/pedrohscramos/teste-frete-rapido/routers"
	"github.com/pedrohscramos/teste-frete-rapido/utils"
)

// @title Frete Rápido API
// @version 1.0
// @description Esta é uma API do Frete Rápido em Go
// @BasePath /
func main() {
	//Configuração
	err := config.Load()
	utils.Error(err, nil)

	//Conexão SQLServer
	dbSql, err := mysql.Connect()
	utils.Error(err, nil)

	router := routers.SetupRoute(dbSql)

	if err := http.ListenAndServe(config.ServerConfig(), router); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

}
