package models

type Breed struct {
	ID         uint `gorm:"primarykey"`
	Name       string
	CategoryID uint
	Category   Category `gorm:"references:ID"`
}
