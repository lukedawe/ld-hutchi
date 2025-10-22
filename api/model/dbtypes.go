package model

type Category struct {
	ID     uint `gorm:"primarykey"`
	Name   string
	Breeds []Breed
}

type Breed struct {
	ID         uint `gorm:"primarykey"`
	Name       string
	CategoryID uint
}
