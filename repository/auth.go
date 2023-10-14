package repository

import (
	"tugas8/connect"
	"tugas8/models"
)

func Register(user *models.User) error {
	result := config.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
