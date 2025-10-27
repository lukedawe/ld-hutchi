package requests

import (
	"errors"
	error_responses "lukedawe/hutchi/handlers/dtos/responses/errors"
	"regexp"
)

const (
	categoryNameMaxLength = 20
)

// Name cannot:
// - have whitespace,
// - be empty,
// - contain characters that are not in a-z
// - should be less than 20 chars
var safeNameRegexCategory = regexp.MustCompile(`^[a-z]*$`)

func ValidateCategoryName(name string) error {
	// Make sure the name isn't now empty string
	if name == "" {
		return errors.New(emptyNameErrorMessage)
	}

	if !safeNameRegexCategory.MatchString(name) {
		return errors.New(invalidStringErrorFormatted(name))
	}

	if len(name) > categoryNameMaxLength {
		return errors.New(nameTooLongErrorFormatted(name))
	}

	return nil
}

func (input *CategoryNameRequiredJson) ValidateCategoryNameRequiredJson() error {
	return ValidateCategoryName(input.Name)
}

func (request *AddCategories) Validate() error {
	// Validate each category individually.
	var err error
	for _, category := range request.Categories {
		err = errors.Join(category.Validate())
	}
	if err != nil {
		return err
	}

	return nil
}

// Validates the add category request and returns nil or response-ready error.
func (request *AddCategory) Validate() error {
	if err := request.ValidateCategoryNameRequiredJson(); err != nil {
		response := error_responses.ErrBadRequestInvalidParam.SetError(err)
		response.Message = err.Error() // Copy error message because it's user facing.
		return response
	}

	var err error
	for _, breed := range request.Breeds {
		err = errors.Join(breed.ValidateBreedNameRequiredJson())
	}
	if err != nil {
		response := error_responses.ErrBadRequestInvalidParam.SetError(err)
		response.Message = err.Error() // Copy error message because it's user facing.
		return response
	}

	return nil
}

func (body *PutCategoryBody) Validate() error {
	if err := body.ValidateCategoryNameRequiredJson(); err != nil {
		response := error_responses.ErrBadRequestInvalidParam.SetError(err)
		response.Message = err.Error() // Copy error message because it's user facing.
		return response
	}

	var err error
	for _, breed := range body.Breeds {
		err = errors.Join(breed.ValidateBreedNameRequiredJson())
	}
	if err != nil {
		response := error_responses.ErrBadRequestInvalidParam.SetError(err)
		response.Message = err.Error() // Copy error message because it's user facing.
		return response
	}
	return nil
}
