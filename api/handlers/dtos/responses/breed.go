package responses

type BreedCreated struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	CategoryId uint   `json:"category_id"`
}
