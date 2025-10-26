package validation

import (
	"errors"
	"regexp"
)

var safeNameRegex = regexp.MustCompile(`^[a-z]*$`)

// Name cannot:
// - have whitespace,
// - be empty,
// - contain characters that are not in a-z
// - should be less than 20 chars
func ValidateName(name string) error {
	// Make sure the name isn't now empty string
	if name == "" {
		return errors.New(emptyNameErrorMessage)
	}

	if !safeNameRegex.MatchString(name) {
		return errors.New(invalidStringErrorFormatted(name))
	}

	if len(name) > 20 {
		return errors.New(nameTooLongErrorFormatted(name))
	}

	return nil
}
