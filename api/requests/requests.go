package requests

import (
	"lukedawe/hutchi/dtos"
)

// Separate out the message requests from the data that we are adding.
type AddCategoryMessage struct {
	Category dtos.Category
}
