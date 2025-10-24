package requests

type AddCategory struct {
	Name   string   `json:"name" binding:"required"`
	Breeds []string `json:"breeds" binding:"required"`
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
