package errors

// For use with custom error-handling middleware.
type ErrorResponse struct {
	Code     string `json:"code"`          // Must be from `/errors/db_errors.go`.
	Status   int    `json:"-"`             // Should be from http package.
	Message  string `json:"message"`       // User-facing error message.
	DebugErr string `json:"debug_message"` // Only shown in debug mode.
}

func (err ErrorResponse) Error() string {
	return err.Message
}

// Omit the Debug message.
func (err ErrorResponse) ToProductionErrorStruct() ErrorResponse {
	return ErrorResponse{
		Code:    err.Code,
		Status:  err.Status,
		Message: err.Message,
	}
}

// Set the internal error for debugging purposes.
func (resp ErrorResponse) SetError(err error) ErrorResponse {
	resp.DebugErr = err.Error()
	return resp
}
