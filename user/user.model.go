package user

import (
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Name  string
	Email string
}
