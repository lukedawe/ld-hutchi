package responses

type BreedWithCategory struct {
	Name     string `json:"name" binding:"required"`
	Category string `json:"category" binding:"required"`
}
