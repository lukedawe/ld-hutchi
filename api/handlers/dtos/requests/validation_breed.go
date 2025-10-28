package requests

import (
	"errors"
	error_responses "lukedawe/hutchi/handlers/dtos/responses/errors"
	"regexp"
)

const (
	breedNameMaxLength = 20
)

// Name cannot:
// - have whitespace,
// - be empty,
// - contain characters that are not in a-z
// - should be less than 20 chars
var safeNameRegexBreed = regexp.MustCompile(`^[a-z]*$`)

func ValidateBreedName(name string) error {
	// Make sure the name isn't now empty string
	if name == "" {
		return errors.New(emptyNameErrorMessage)
	}

	if !safeNameRegexBreed.MatchString(name) {
		return errors.New(invalidStringErrorFormatted(name))
	}

	if len(name) > breedNameMaxLength {
		return errors.New(nameTooLongErrorFormatted(name))
	}

	return nil
}

func (input *BreedNameRequiredJson) ValidateBreedNameRequiredJson() error {
	return ValidateBreedName(input.Name)
}

// Validates the `AddBreed` struct and returns a user-facing response if it's not valid.
func (breed *AddBreed) Validate() error {
	if err := breed.ValidateBreedNameRequiredJson(); err != nil {
		response := error_responses.ErrBadRequestValidation.SetError(err)
		response.Message = err.Error()
		return response
	}
	return nil
}

func (breed *PatchBreedBody) Validate() error {
	if err := breed.ValidateBreedNameRequiredJson(); err != nil {
		response := error_responses.ErrBadRequestValidation.SetError(err)
		response.Message = err.Error()
		return response
	}
	return nil
}
