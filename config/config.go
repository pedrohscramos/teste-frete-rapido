package config

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API   APIConfig
	DBSql DBConfig
}

type APIConfig struct {
	Port          string
	Documentation string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load() error {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}

	cfg = new(config)
	cfg.API = APIConfig{
		Port:          viper.GetString("api.port"),
		Documentation: viper.GetString("api.documentation"),
	}

	cfg.DBSql = DBConfig{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetString("mysql.port"),
		User:     viper.GetString("mysql.user"),
		Password: viper.GetString("mysql.password"),
		Database: viper.GetString("mysql.database"),
	}

	return nil
}

func GetAPIConfig() APIConfig {
	return cfg.API
}

func GetDBSql() DBConfig {
	return cfg.DBSql
}
