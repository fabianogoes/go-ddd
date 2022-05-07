package repository

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

const (
	maxAttempts = 20
)

// BaseDBO represents a data entity base
type BaseDBO struct {
	ID         uuid.UUID `gorm:"column:id;type:uuid;primarykey"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	ModifiedAt time.Time `gorm:"column:modified_at"`
}

func NewPostgresDB(config *Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		db, err = gorm.Open(postgres.Open(config.dsn()), &gorm.Config{})
		if err == nil {
			break
		}
		log.Warnf("postgres connection failed - %v", err)
		time.Sleep(time.Duration(attempts) * time.Second)
	}
	if err != nil {
		return nil, err
	}
	log.Infoln("postgres connected")

	return db, err
}

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// dns returns connection string
func (c *Config) dsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
		c.SSLMode)
}
