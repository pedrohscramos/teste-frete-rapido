package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ServerConfiguration struct {
	Port            string
	Secret          string
	LimitPerRequest int
}

func ServerConfig() string {
	server := fmt.Sprintf("%s:%s", viper.GetString("api.host"), viper.GetString("api.port"))
	log.Print("Servidor rodando em: ", server)
	return server
}
