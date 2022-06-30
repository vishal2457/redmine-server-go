package models

import (
	"time"
)

// User is the main user model.
type ProjectUser struct {
	ID        uint      `json:"id" gorm:"->;primaryKey"`
	UserID    uint      `json:"user_id"`
	ProjectID uint      `json:"project_id"`
	CreatedAt time.Time `json:"created_at"`
}
