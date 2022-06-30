package models

import (
	"time"

	"gorm.io/gorm"
)

// User is the main user model.
type Project struct {
	gorm.Model
	Name         string    `json:"name"`
	LiveLink     string    `json:"liveLink"`
	ApiLink      string    `json:"api_link"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ProjectUsers []ProjectUser
	Tasks        []Task
}

// GetUsers queries the database for all users.
func GetProjects(projects *[]Project) (err error) {
	if err = DB.Find(projects).Error; err != nil {
		return err
	}

	return nil
}
