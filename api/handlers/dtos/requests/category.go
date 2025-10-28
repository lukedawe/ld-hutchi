package requests

type CategoryNameRequiredJson struct {
	Name string `json:"name" binding:"required"`
}

type CategoryNameRequiredUri struct {
	Name string `uri:"name" binding:"required"`
}

type CategoryIdRequiredUri struct {
	Id uint `uri:"id" binding:"required"`
}

type CategoryIdOptionalUri struct {
	Id uint `uri:"id"`
}

type Breed struct {
	BreedNameRequiredJson
}

type BreedArrayRequired struct {
	Breeds []Breed `json:"breeds" binding:"required"`
}

type CategoryArrayRequired struct {
	Categories []AddCategory `json:"categories"`
}

type AddCategory struct {
	CategoryNameRequiredJson
	BreedArrayRequired
}

type AddCategories struct {
	CategoryArrayRequired
}

type GetCategoriesToBreeds struct {
	paginated
}

type GetCategory struct {
	CategoryIdRequiredUri
}

type GetCategoryToBreeds struct {
	CategoryIdRequiredUri
}

type PutCategoryUri struct {
	CategoryIdOptionalUri
}

type PutCategoryBody struct {
	CategoryNameRequiredJson
	BreedArrayRequired
}
