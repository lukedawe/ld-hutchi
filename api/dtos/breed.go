package dtos

type Breed struct {
	Name string `json:"name" binding:"required"`
}
