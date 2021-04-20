package database

import (
	"fmt"
	. "github.com/gguibittencourt/go-restapi/config"
	"github.com/gguibittencourt/go-restapi/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(New)
var config = Config{}

type Params struct {
	fx.In

	Logger *zap.Logger
}

func New(p Params) (*gorm.DB, error) {
	config.Read()
	db, err := connectDB(&config)
	if err != nil {
		p.Logger.Info(fmt.Sprintf("Error connecting to database : error=%v\n",  err))
		return nil, err
	}
	p.Logger.Info("Connected to database")
	db.AutoMigrate(&models.User{})
	return db, nil
}

func connectDB(config *Config) (*gorm.DB, error) {
	dbSource := fmt.Sprintf("%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", config.Username, config.Server, config.Database)
	return gorm.Open("mysql", dbSource)
}
