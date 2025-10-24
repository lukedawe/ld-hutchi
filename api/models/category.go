package models

import "lukedawe/hutchi/dtos"

type Category struct {
	ID     uint    `gorm:"primarykey"`
	Name   string  `gorm:"unique"`
	Breeds []Breed `gorm:"foreignKey:CategoryID"`
}

func (c *Category) ToResponse() dtos.Category {
	breeds := BreedsToDTO(c.Breeds)

	return dtos.Category{
		Name:   c.Name,
		Breeds: breeds,
	}
}

func CategoryDtoToModel(c *dtos.Category) Category {
	breeds := make([]Breed, len(c.Breeds))

	for i, breed := range c.Breeds {
		breeds[i] = BreedDtoToModel(&breed)
	}

	return (Category{
		Name:   c.Name,
		Breeds: breeds,
	})
}

func CategoriesToDTO(categories []Category) []dtos.Category {
	response := make([]dtos.Category, len(categories))

	for i, category := range categories {
		response[i] = category.ToResponse()
	}

	return response
}
