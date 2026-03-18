package repos

import (
	"database/sql"
	"github.com/kungrem23/quizgo/internal/store/models"
	"log"
)

type ImageRepo struct {
	db *sql.DB
}

func NewImageRepo(db *sql.DB) *ImageRepo {
	return &ImageRepo{db: db}
}

func (r *ImageRepo) CreateNewImage(imageURL string) (*models.Image, error) {
	query := `INSERT INTO images
	(image_url)
	VALUES ($1)
	RETURNING id, image_url;`
	row := r.db.QueryRow(query, imageURL)
	image := models.NewImage()
	err := row.Scan(&image.Id, &image.ImageURL)
	if err != nil {
		log.Printf("Adding Image error: %v\n", err)
		return nil, err
	}
	return image, nil
}

func (r *ImageRepo) DeleteImage(id int) error {
	query := `DELETE FROM images WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Deleting image(id=%v) error: %v", id, err)
		return err
	}
	return nil
}

func (r *ImageRepo) GetImageById(id int) (*models.Image, error) {
	query := `SELECT id, image_url FROM images WHERE id=$1`
	row := r.db.QueryRow(query, id)
	image := models.NewImage()
	err := row.Scan(&image.Id, &image.ImageURL)
	if err != nil {
		log.Printf("Get image(id=%v) error: %v", id, err)
		return nil, err
	}
	return image, nil
}

// func (r *ImageRepo) GetAllImages() ([]*store.Image, error) {

// }
