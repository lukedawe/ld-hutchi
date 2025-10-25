package services

import (
	"errors"
	"log"
	lde "lukedawe/hutchi/dtos/responses/errors"

	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// Inspect the raw GORM error and translates it into a specific, public-facing ErrorResponse.
func TranslateDbError(err error) lde.ErrorResponse {
	// Check for standard GORM errors first
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return lde.ErrNotFound.SetError(err)
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return lde.ErrConflictUniqueName.SetError(err)
	}
	// Check for context cancellation/timeout errors (common GORM/DB pattern)
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return lde.ErrInternalUnknown.SetError(err)
	}

	var pgErr *pgconn.PgError
	// Map specific PostgreSQL SQLSTATE codes (pgErr.Code) to your API errors
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // unique_violation (e.g., trying to add a Category name that exists)
			return lde.ErrConflictUniqueName.SetError(err)

		case "23503": // foreign_key_violation (e.g., linking Breed to non-existent CategoryID)
			return lde.ErrNotFound.SetError(err)

		case "23000": // integrity_constraint_violation (General integrity failure)
			return lde.ErrBadRequestValidation.SetError(err)

		case "22P02": // invalid_text_representation (e.g., non-integer passed to an integer column)
			return lde.ErrBadRequestBinding.SetError(err)

		default:
			log.Println("Error code not recognised: ", pgErr.Code)
		}

	}

	// Default to an Unknown Error if the cause is unhandled
	return lde.ErrInternalUnknown.SetError(err)
}
