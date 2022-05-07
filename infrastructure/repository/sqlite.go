package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteDB(config *Config) (*gorm.DB, error) {
	var db *gorm.DB
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Infoln("sqlite connected")

	return db, err
}
