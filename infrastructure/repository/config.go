package repository

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// BaseDBO represents a data entity base
type BaseDBO struct {
	ID         uuid.UUID `gorm:"column:id;type:uuid;primarykey"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	ModifiedAt time.Time `gorm:"column:modified_at"`
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
func (c *Config) dsnPostgres() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
		c.SSLMode)
}

func NewPostgresConfigTest() *Config {
	return &Config{
		Host:     "localhost",
		Port:     5432,
		User:     "test",
		Password: "test",
		DBName:   "test",
		SSLMode:  "disable",
	}
}

func NewSqliteConfigTest() *Config {
	return &Config{}
}
