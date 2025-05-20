package product

import (
	"go-crud/user"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Price int

	CreatedBy   uint      // Foreign key
	CreatedUser user.User `gorm:"foreignKey:CreatedBy;references:ID"` // Relasi ke user
}
