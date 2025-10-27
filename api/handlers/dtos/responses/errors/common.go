package errors

import "net/http"

var (
	// 500 Internal Server Errors
	ErrInternalUnknown = ErrorResponse{
		Code:    UnknownErrorCode,
		Status:  http.StatusInternalServerError,
		Message: UnknownErrorMessage,
	}
	ErrInternalDatabase = ErrorResponse{
		Code:    DatabaseErrorCode,
		Status:  http.StatusInternalServerError,
		Message: DatabaseErrorMessage,
	}
	ErrInternalExternalService = ErrorResponse{
		Code:    ExternalServiceFailure,
		Status:  http.StatusServiceUnavailable, // Use 503 for service failure
		Message: ExternalServiceFailureMessage,
	}

	// 400 Bad Request Errors
	ErrBadRequestBinding = ErrorResponse{
		Code:    BindingErrorCode,
		Status:  http.StatusBadRequest,
		Message: BindingErrorMessage,
	}
	ErrBadRequestValidation = ErrorResponse{
		Code:    ValidationErrorCode,
		Status:  http.StatusBadRequest,
		Message: ValidationErrorMessage,
	}
	ErrBadRequestRequiredField = ErrorResponse{
		Code:    RequiredFieldMissing,
		Status:  http.StatusBadRequest,
		Message: RequiredFieldMissingMessage,
	}
	ErrBadRequestInvalidParam = ErrorResponse{
		Code:    InvalidParameterCode,
		Status:  http.StatusBadRequest,
		Message: InvalidParameterMessage,
	}
	ErrBadRequestInvalidJSON = ErrorResponse{
		Code:    InvalidJsonCode,
		Status:  http.StatusBadRequest,
		Message: InvalidJsonMessage,
	}

	// 401/403 Authorization Errors
	ErrUnauthorized = ErrorResponse{
		Code:    UnauthenticatedCode,
		Status:  http.StatusUnauthorized,
		Message: UnauthenticatedMessage,
	}
	ErrUnauthorizedInvalidCredentials = ErrorResponse{
		Code:    InvalidCredentialsCode,
		Status:  http.StatusUnauthorized,
		Message: InvalidCredentialsMessage,
	}
	ErrForbidden = ErrorResponse{
		Code:    ForbiddenCode,
		Status:  http.StatusForbidden,
		Message: ForbiddenMessage,
	}
	ErrForbiddenAccountInactive = ErrorResponse{
		Code:    AccountInactiveCode,
		Status:  http.StatusForbidden,
		Message: AccountInactiveMessage,
	}

	// 404 Not Found
	ErrNotFound = ErrorResponse{
		Code:    ResourceNotFoundCode,
		Status:  http.StatusNotFound,
		Message: ResourceNotFoundMessage,
	}

	// 405 Method Not Allowed
	ErrMethodNotAllowed = ErrorResponse{
		Code:    MethodNotAllowedCode,
		Status:  http.StatusMethodNotAllowed,
		Message: MethodNotAllowedMessage,
	}

	// 409 Conflict Errors
	ErrConflictUniqueName = ErrorResponse{
		Code:    UniqueNameConflictCode,
		Status:  http.StatusConflict,
		Message: UniqueNameConflictMessage,
	}
	ErrConflictState = ErrorResponse{
		Code:    StateConflictCode,
		Status:  http.StatusConflict,
		Message: StateConflictMessage,
	}
	ErrConflictConcurrentUpdate = ErrorResponse{
		Code:    ConcurrentUpdateCode,
		Status:  http.StatusConflict,
		Message: ConcurrentUpdateMessage,
	}

	// 415 Unsupported Media Type
	ErrUnsupportedMediaType = ErrorResponse{
		Code:    UnsupportedMediaType,
		Status:  http.StatusUnsupportedMediaType,
		Message: UnsupportedMediaTypeMessage,
	}

	// 429 Rate Limit
	ErrRateLimitExceeded = ErrorResponse{
		Code:    RateLimitExceededCode,
		Status:  http.StatusTooManyRequests,
		Message: RateLimitExceededMessage,
	}
)
