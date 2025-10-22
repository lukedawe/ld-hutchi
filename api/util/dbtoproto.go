package util

import (
	"lukedawe/hutchi/generated/proto_dogs"
	"lukedawe/hutchi/model"
)

func CategoriesToAllDogs(categories []model.Category) *proto_dogs.AllDogs {
	// Encode into a protobufs response
	allDogs := &proto_dogs.AllDogs{}
	for _, category := range categories {
		breedNames := make([]string, len(category.Breeds))
		for _, breed := range category.Breeds {
			breedNames = append(breedNames, breed.Name)
		}
		allDogs.Categories = append(allDogs.Categories, &proto_dogs.DogCategory{Name: category.Name, Breeds: breedNames})
	}
	return allDogs
}
