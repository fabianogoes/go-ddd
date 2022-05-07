package postgres

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-ddd/infrastructure/repository/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(config *database.Config) (*gorm.DB, error) {
	var db *gorm.DB
	db, err := gorm.Open(postgres.Open(dsn(config)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Infoln("postgres connected")

	return db, err
}

func dsn(c *database.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
		c.SSLMode)
}
