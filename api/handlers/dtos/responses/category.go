package responses

type CategoryBreed struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type CategoryCreated struct {
	Id     uint            `json:"id"`
	Name   string          `json:"name"`
	Breeds []CategoryBreed `json:"breeds"`
}

type CategoriesCreated struct {
	Categories []CategoryCreated `json:"categories"`
}
