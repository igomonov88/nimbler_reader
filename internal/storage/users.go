package storage

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"go.opencensus.io/trace"
)

var (
	// ErrNotFound is used when a specific User is requested but does not exist.
	ErrNotFound = errors.New("User not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidUserID = errors.New("ID is not in its proper form")
)

// UpdateUserInfo replaces a user document in the database.
func GetUserInfo(ctx context.Context, db *sqlx.DB, userID string) (*User, error) {
	ctx, span := trace.StartSpan(ctx, "internal.user.Update")
	defer span.End()

	if _, err := uuid.Parse(userID); err != nil {
		return nil, ErrInvalidUserID
	}

	const q = `SELECT * FROM users WHERE user_id = $1`
	var u User

	if err := db.GetContext(ctx, &u, q, userID); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, constraintError(err)
	}

	return &u, nil
}

func constraintError(err error) error {
	const (
		foreignKeyViolationCode = "23503"
	)
	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Code == foreignKeyViolationCode {
			return ErrInvalidUserID
		}
	}
	return err
}
