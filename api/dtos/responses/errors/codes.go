package errors

// /internal/api/errorcodes.go
const (
	// Generic/5xx
	UnknownErrorCode       = "unknown_error"
	DatabaseErrorCode      = "database_error"
	ExternalServiceFailure = "external_service_failure"

	// 400 Bad Request
	BindingErrorCode     = "binding_error"
	ValidationErrorCode  = "validation_error"
	RequiredFieldMissing = "required_field_missing"
	InvalidParameterCode = "invalid_parameter"
	InvalidJsonCode      = "invalid_json"
	UnsupportedMediaType = "unsupported_media_type"

	// 401/403 Auth/Permission
	UnauthenticatedCode    = "unauthenticated"
	InvalidCredentialsCode = "invalid_credentials"
	ForbiddenCode          = "forbidden"
	AccountInactiveCode    = "account_inactive"

	// 404 Not Found
	ResourceNotFoundCode = "resource_not_found"

	// 409 Conflict
	UniqueNameConflictCode = "unique_name_conflict"
	StateConflictCode      = "state_conflict"
	ConcurrentUpdateCode   = "concurrent_update"

	// 429 Rate Limit
	RateLimitExceededCode = "rate_limit_exceeded"

	// 405 Method
	MethodNotAllowedCode = "method_not_allowed"
)
