package requests

type paginated struct {
	Page     uint `uri:"page" binding:"required"`
	PageSize uint `uri:"page_size" binding:"required"`
}
