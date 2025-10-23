package dtos

type Category struct {
	Name   string  `binding:"required"`
	Breeds []Breed `binding:"required"`
}
