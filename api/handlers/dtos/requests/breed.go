package requests

// Any json body that requires a breed name.
// This is for use with validation.
type BreedNameRequiredJson struct {
	Name string `json:"name" binding:"required"`
}

// Same for Uri.
type BreedNameRequiredUri struct {
	Name string `uri:"name" binding:"required"`
}

type BreedCategoryNameRequiredJson struct {
	// Avoid circular references.
	CategoryName string `json:"category_name" binding:"required"`
}

type AddBreed struct {
	BreedNameRequiredJson
	BreedCategoryNameRequiredJson
}

type GetBreed struct {
	BreedNameRequiredUri
}
