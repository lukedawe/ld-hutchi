package responses

type CategoryBreed struct {
	Name string
}

type CategoryCreated struct {
	Name   string          `json:"name" binding:"required"`
	Breeds []CategoryBreed `json:"breeds" binding:"required"`
}

type CategoriesCreated struct {
	Categories []CategoryCreated `json:"categories"`
}
