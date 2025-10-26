package models

type Category struct {
	ID     uint    `gorm:"primarykey"`
	Name   string  `gorm:"unique"`
	Breeds []Breed `gorm:"foreignKey:CategoryID"`
}
