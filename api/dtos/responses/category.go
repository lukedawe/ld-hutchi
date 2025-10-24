package responses

type Category struct {
	Name   string  `json:"name" binding:"required"`
	Breeds []Breed `json:"breeds" binding:"required"`
}
