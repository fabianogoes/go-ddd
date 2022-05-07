package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(config *Config) (*gorm.DB, error) {
	var db *gorm.DB
	db, err := gorm.Open(postgres.Open(config.dsnPostgres()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Infoln("postgres connected")

	return db, err
}
