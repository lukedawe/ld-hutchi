package main

type Category struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

type Breed struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	Category uint
}
