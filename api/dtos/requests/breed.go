package requests

type AddBreed struct {
	Name         string `json:"name" binding:"required"`
	CategoryName string `json:"category_name" binding:"required"`
}

type GetBreed struct {
	Name string `uri:"name" binding:"required"`
}
