package database

import (
	"fmt"
	. "github.com/gguibittencourt/go-restapi/config"
	"github.com/gguibittencourt/go-restapi/models"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	_ = db.AutoMigrate(&models.Task{})
	return db, nil
}

func connectDB(config *Config) (*gorm.DB, error) {
	dbSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", config.Username, config.Password, config.Server, config.Database)
	return gorm.Open(mysql.Open(dbSource), &gorm.Config{})
}
