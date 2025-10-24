package models

import (
	"lukedawe/hutchi/dtos"
)

type Breed struct {
	ID         uint `gorm:"primarykey"`
	Name       string
	CategoryID uint
	Category   Category `gorm:"references:ID"`
}

func (b *Breed) ToResponse() dtos.Breed {
	return dtos.Breed{
		Name: b.Name,
	}
}

func (b *Breed) ToResponseWithCategory() dtos.BreedWithCategory {
	return dtos.BreedWithCategory{
		Name:     b.Name,
		Category: b.Category.Name,
	}
}

func BreedDtoToModel(b *dtos.Breed) Breed {
	return Breed{
		Name: b.Name,
	}
}

func BreedsToDTO(breeds []Breed) []dtos.Breed {
	response := make([]dtos.Breed, len(breeds))

	for i, breed := range breeds {
		response[i] = breed.ToResponse()
	}

	return response
}

func BreedsToDtoWithCategory(breeds []Breed) []dtos.BreedWithCategory {
	response := make([]dtos.BreedWithCategory, len(breeds))

	for i, breed := range breeds {
		response[i] = breed.ToResponseWithCategory()
	}

	return response
}
