package storage_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/igomonov88/nimbler_reader/internal/storage"
	"github.com/igomonov88/nimbler_reader/internal/tests"
)

func TestURLs(t *testing.T) {
	db, teardown := tests.NewUnit(t)
	defer teardown()

	t.Log("Given the need to test url functionality:")

	url, err := storage.RetrieveOriginalURL(context.Background(), db, "4ds35d")
	if err != nil {
		t.Fatalf("\t%s\tShould be able to get original url from storage: %s", tests.Failed, err)
	}

	if !cmp.Equal(url, "https://tut.by") {
		t.Fatalf("\t%s\tShould return original url: %s", tests.Failed, err)
	}

	t.Logf("\t%s\tShould be able to get original url from storage.", tests.Success)
}
