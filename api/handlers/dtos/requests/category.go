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

type AddCategoryJson struct {
	CategoryNameRequiredJson
	BreedArrayRequired
}

type AddCategories struct {
	Categories []AddCategoryJson `json:"categories"`
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

type PatchCategoryUri struct {
	CategoryIdRequiredUri
}

type PatchCategoryBody struct {
	CategoryNameRequiredJson
}
type DeleteCategoryUri struct {
	CategoryIdRequiredUri
}
