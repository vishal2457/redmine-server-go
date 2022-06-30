package models

import (
	"time"

	"gorm.io/gorm"
)

// User is the main user model.
type Task struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"->;primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ProjectID   uint      `json:"project"`
	CreatedBy   uint      `json:"user_by"`
	UserID      uint      `json:"user_id" gorm:"default:null"`
	Type        uint      `json:"type"`
	DueDate     time.Time `json:"due_date"`
	Priority    string    `json:"priority"`
	Hours       uint      `json:"hours"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func CreateTask(task *Task) (err error) {
	if err = DB.Create(task).Error; err != nil {
		return err
	}

	return nil
}

// GetUsers queries the database for all users.
func GetTasks(tasks *[]Task) (err error) {
	if err = DB.Find(tasks).Error; err != nil {
		return err
	}

	return nil
}
