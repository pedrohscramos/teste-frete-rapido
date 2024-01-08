package mysql

import (
	"github.com/pedrohscramos/teste-frete-rapido/config"
	"github.com/pedrohscramos/teste-frete-rapido/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Database *gorm.DB
}

func Connect() (*DB, error) {
	conf := config.GetDBSql()
	db, err := gorm.Open(mysql.Open(conf.User+":"+conf.Password+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.Database+"?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	utils.Error(err, nil)

	return &DB{Database: db}, nil
}
