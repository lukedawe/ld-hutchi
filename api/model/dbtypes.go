package model

type Category struct {
	ID     uint `gorm:"primarykey" json:"-"`
	Name   string
	Breeds []Breed
}

type Breed struct {
	ID         uint `gorm:"primarykey" json:"-"`
	Name       string
	CategoryID uint
}
