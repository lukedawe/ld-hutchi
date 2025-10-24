package responses

type Breed struct {
	Name string `json:"name" binding:"required"`
}
