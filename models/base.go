package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/google/uuid"
)

// Base contains common columns for all tables.
type Base struct {
 ID        uuid.UUID `json:"id";gorm:"type:uuid;primary_key;"`
 CreatedAt time.Time
 UpdatedAt time.Time
 DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
 uuid := uuid.New()
 return scope.SetColumn("ID", uuid)
}