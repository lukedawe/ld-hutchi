package responses

type CategoryBreed struct {
	Name string
}

type Category struct {
	Name   string          `json:"name" binding:"required"`
	Breeds []CategoryBreed `json:"breeds" binding:"required"`
}

type Categories struct {
	Categories []Category `json:"categories"`
}
