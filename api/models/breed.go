package models

import "lukedawe/hutchi/dtos"

type Breed struct {
	ID         uint `gorm:"primarykey"`
	Name       string
	CategoryID uint `json:"-"`
}

func (b *Breed) ToResponse() dtos.Breed {
	return dtos.Breed{
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
