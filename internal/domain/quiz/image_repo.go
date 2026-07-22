package quiz

import "context"

// "database/sql"
// "github.com/kungrem23/quizgo/internal/store/models"
// "log"

// type ImageRepo struct {
// 	db *sql.DB
// }

// func NewImageRepo(db *sql.DB) *ImageRepo {
// 	return &ImageRepo{db: db}
// }

func (r *PostgresRepository) CreateNewImage(ctx context.Context, imageURL string) error {
	query := `INSERT INTO images
	(image_url)
	VALUES ($1)
	RETURNING id, image_url;`
	_, err := r.db.ExecContext(ctx, query, imageURL)
	// image := NewImage()
	// err := row.Scan(&image.Id, &image.ImageURL)
	// if err != nil {
	// 	// log.Printf("Adding Image error: %v\n", err)
	// 	return nil, err
	// }
	return err
}

func (r *PostgresRepository) DeleteImage(ctx context.Context, id string) error {
	query := `DELETE FROM images WHERE id = $1`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNotFound
	}
	// if err != nil {
	// 	// log.Printf("Deleting image(id=%v) error: %v", id, err)
	// 	return err
	// }
	return nil
}

func (r *PostgresRepository) GetImage(ctx context.Context, id string) (Image, error) {
	query := `SELECT id, image_url FROM images WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	var image Image
	err := row.Scan(&image.Id, &image.ImageURL)
	// if err != nil {
	// 	// log.Printf("Get image(id=%v) error: %v", id, err)
	// 	return image, err
	// }
	return image, err
}
