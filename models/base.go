package models

import (
	"time"
)

// Model db base model
type Model struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
