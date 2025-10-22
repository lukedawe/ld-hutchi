package handlers

import "gorm.io/gorm"

// Using struct embedding pattern.
type Handler struct {
	DB *gorm.DB // The database connection or pool
}
