package models

import (
	"github/Hiinnn/practice-go/config"
)

// GetAllUsers -> Get all user data
func GetAllUsers(user *[]User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser -> Create new user
func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
