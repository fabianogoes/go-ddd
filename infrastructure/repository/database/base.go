package database

import (
	"github.com/google/uuid"
	"time"
)

type BaseDBO struct {
	ID         uuid.UUID `gorm:"column:id;type:uuid;primarykey"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	ModifiedAt time.Time `gorm:"column:modified_at"`
}
