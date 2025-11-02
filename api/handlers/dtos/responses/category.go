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

type CategoryResponsePaginaged struct {
	Categories []CategoryCreated `json:"categories"`
	PageSize   uint              `json:"page_size"`
	NoPages    int64             `json:"no_pages"`
}

type CategoriesCreated struct {
	Categories []CategoryCreated `json:"categories"`
}
