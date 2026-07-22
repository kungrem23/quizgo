package quiz

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// ============== CreateNewImage ==============

func TestCreateNewImage_Success(t *testing.T) {
	repo, mock, cleanup := newPostgresRepoMock(t)
	defer cleanup()
	const (
		imageURL = "bebra.com"
	)
	query := regexp.QuoteMeta(`INSERT INTO images
	(image_url)
	VALUES ($1)
	RETURNING id, image_url;`)
	mock.ExpectExec(query).
		WithArgs(imageURL).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := repo.CreateNewImage(context.Background(), imageURL)
	if err != nil {
		t.Fatalf("CreateNewImage error got: %v | expect: %v",
			err, nil)
	}
}

// =============== DeleteImage ================

func TestDeleteImage_Success(t *testing.T) {
	repo, mock, cleanup := newPostgresRepoMock(t)
	defer cleanup()
	const (
		imageId = "qwert"
	)
	query := regexp.QuoteMeta(`DELETE FROM images WHERE id = $1`)
	mock.ExpectExec(query).
		WithArgs(imageId).
		WillReturnResult(sqlmock.NewResult(0, 1))
	err := repo.DeleteImage(context.Background(), imageId)
	if err != nil {
		t.Fatalf("DeleteImage error got: %v | expect: %v",
			err, nil)
	}
}

func TestDeleteImage_NotFound(t *testing.T) {
	repo, mock, cleanup := newPostgresRepoMock(t)
	defer cleanup()
	const (
		imageId = "qwert"
	)
	query := regexp.QuoteMeta(`DELETE FROM images WHERE id = $1`)
	mock.ExpectExec(query).
		WithArgs(imageId).
		WillReturnResult(sqlmock.NewResult(0, 0))
	err := repo.DeleteImage(context.Background(), imageId)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("DeleteImage error got: %v | expect: %v",
			err, ErrNotFound)
	}
}

// ================ GetImage ==================

func TestGetImage_Success(t *testing.T) {
	repo, mock, cleanup := newPostgresRepoMock(t)
	defer cleanup()
	const (
		imageId  = "qwert"
		imageURL = "bebra.com"
	)
	query := regexp.QuoteMeta(`SELECT id, image_url FROM images WHERE id=$1`)
	mock.ExpectQuery(query).
		WithArgs(imageId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "image_url"}).
			AddRow(imageId, imageURL))
	want := Image{Id: imageId, ImageURL: imageURL}
	got, err := repo.GetImage(context.Background(), imageId)
	if err != nil {
		t.Fatalf("DeleteImage error got: %v | expect: %v",
			err, nil)
	}
	if got != want {
		t.Fatalf("DeleteImage got: %#v | expect: %#v",
			err, nil)
	}
}

func TestGetImage_NotFound(t *testing.T) {
	repo, mock, cleanup := newPostgresRepoMock(t)
	defer cleanup()
	const (
		imageId  = "qwert"
		imageURL = "bebra.com"
	)
	query := regexp.QuoteMeta(`SELECT id, image_url FROM images WHERE id=$1`)
	mock.ExpectQuery(query).
		WithArgs(imageId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "image_url"}))
	_, err := repo.GetImage(context.Background(), imageId)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("DeleteImage error got: %v | expect: %v",
			err, sql.ErrNoRows)
	}
}
