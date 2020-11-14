package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.opencensus.io/trace"
)

// ErrURLNotFound is used when url is requested by the key does not exist.
var ErrURLNotFound = errors.New("Url not found")

// RetrieveOriginalURL returned original urlPath if it matched with provided key, or
// ErrURLNotFound if there is no path for such key.
func RetrieveOriginalURL(ctx context.Context, db *sqlx.DB, key string) (string, error) {
	ctx, span := trace.StartSpan(ctx, "internal.storage.RetrieveOriginalURL")
	defer span.End()

	const query = `SELECT original_url FROM urls WHERE url_hash = $1 AND expired_at < $2`
	var urlPath string

	if err := db.GetContext(ctx, &urlPath, query, key, time.Now()); err != nil {
		switch err {
		case sql.ErrNoRows:
			return "", ErrURLNotFound
		default:
			return "",
				errors.Wrap(err, "selecting from urls")
		}
	}

	return urlPath, nil
}

// RetrieveAllExpiredURLKeysFromDate return all url hashes which expired.
func RetrieveAllExpiredURLKeysFromDate(ctx context.Context, db *sqlx.DB, date time.Time, limit int32) ([]string, error) {
	ctx, span := trace.StartSpan(ctx, "internal.storage.RetrieveAllExpiredURLKeysFromDate")
	defer span.End()

	const query = `SELECT url_hash FROM urls WHERE expired_at < $1 limit $2`
	var urlHashes []string

	if err := db.SelectContext(ctx, &urlHashes, query, date, limit); err != nil {
		return nil, errors.Wrap(err, "selecting expired url hashes")
	}

	return urlHashes, nil
}

// DoesURLExist returns info about existing url in database.
func DoesURLAliasExist(ctx context.Context, db *sqlx.DB, alias string) (bool, error) {
	ctx, span := trace.StartSpan(ctx, "internal.user.DoesURLAliasExist")
	defer span.End()

	var exist bool
	const q = `SELECT EXISTS(SELECT 1 FROM urls WHERE custom_alias = $1);`

	err := db.GetContext(ctx, &exist, q, alias)
	if err != nil {
		return exist, errors.Wrapf(err, "selecting custom alias exist %q", alias)
	}

	return exist, err
}