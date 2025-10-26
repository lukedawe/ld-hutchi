package validation

const (
	emptyNameErrorMessage = "Empty name provided."
)

func invalidStringErrorFormatted(name string) string {
	return "Name `" + name + "`included invalid characters."
}

func nameTooLongErrorFormatted(name string) string {
	return "Name `" + name + "` too long."
}
