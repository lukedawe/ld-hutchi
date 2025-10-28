package requests

// Any json body that requires a breed name.
// This is for use with validation.
type BreedNameRequiredJson struct {
	Name string `json:"name" binding:"required"`
}

// Same for Uri.
type BreedIdRequiredUri struct {
	Id uint `uri:"id" binding:"required"`
}

type BreedIdOptionalUri struct {
	Id uint `uri:"id"`
}

type BreedCategoryIdRequiredJson struct {
	// Avoid circular references.
	CategoryId uint `json:"category_id" binding:"required"`
}

type GetBreed struct {
	BreedIdRequiredUri
}

type AddBreed struct {
	BreedNameRequiredJson
	BreedCategoryIdRequiredJson
}

type PutBreedUri struct {
	BreedIdOptionalUri
}

type PatchBreedUri struct {
	BreedIdOptionalUri
}

type PatchBreedBody struct {
	BreedNameRequiredJson
}
