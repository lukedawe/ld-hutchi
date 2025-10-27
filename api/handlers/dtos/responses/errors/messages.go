package errors

const (
	// Generic/5xx Messages
	UnknownErrorMessage           = "An unexpected server error occurred. Please try again later."
	DatabaseErrorMessage          = "A temporary database issue occurred. We're working on it."
	ExternalServiceFailureMessage = "A required external service is unavailable right now."

	// 400 Bad Request Messages
	BindingErrorMessage         = "The request data format is incorrect or invalid."
	ValidationErrorMessage      = "One or more fields failed validation."
	RequiredFieldMissingMessage = "A mandatory field is missing from the request."
	InvalidParameterMessage     = "One of the request parameters has an invalid value."
	InvalidJsonMessage          = "The request body is not valid JSON."
	UnsupportedMediaTypeMessage = "The content type of the request is not supported. Please use application/json."

	// 401/403 Auth/Permission Messages
	UnauthenticatedMessage    = "Authentication is required to access this resource."
	InvalidCredentialsMessage = "The username or password provided is incorrect."
	ForbiddenMessage          = "You do not have permission to perform this action."
	AccountInactiveMessage    = "Your account is inactive or has been suspended."

	// 404 Not Found Message
	ResourceNotFoundMessage = "The requested resource could not be found."

	// 409 Conflict Messages
	UniqueNameConflictMessage = "The provided name or identifier is already in use and must be unique."
	StateConflictMessage      = "The request conflicts with the current state of the resource."
	ConcurrentUpdateMessage   = "The resource was modified by another request. Please reload and try again."

	// 429 Rate Limit Message
	RateLimitExceededMessage = "Too many requests. Please try again later."

	// 405 Method Message
	MethodNotAllowedMessage = "The requested method is not allowed for this endpoint."
)
