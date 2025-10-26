package requests

type Breed struct {
	Name string `json:"name" binding:"required"`
}

type AddCategory struct {
	Name   string  `json:"name" binding:"required"`
	Breeds []Breed `json:"breeds" binding:"required"`
}

type AddCategories struct {
	Categories []AddCategory `json:"categories"`
}

type GetCategoriesToBreeds struct {
	Page     uint `uri:"page" binding:"required"`
	PageSize uint `uri:"page_size" binding:"required"`
}

type GetCategory struct {
	Name string `uri:"name" binding:"required"`
}

type GetCategoryToBreeds struct {
	Name string `uri:"name" binding:"required"`
}

type PutCategoryUri struct {
	Name string `uri:"name" binding:"required"`
}

type PutCategoryBody struct {
	Breeds []Breed `json:"breeds" binding:"required"`
}
