package repos

import (
	"database/sql"
	// "fmt"
	"log"

	store "github.com/kungrem23/quizgo/internal/store"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateNewUser(username string, passwordHash string) (*store.User, error) {
	query := `INSERT INTO users
	(username, password_hash)
	VALUES ($1, $2)
	RETURNING id, username, password_hash;`
	row := r.db.QueryRow(query, username, passwordHash)
	user := store.NewUser()
	err := row.Scan(&user.Id, &user.Username, &user.PasswordHash)
	if err != nil {
		log.Printf("Adding user error: %v", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) Deleteuser(id int) error {
	query := "DELETE FROM users WHERE id = $1;"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Deleting user error: %v", err)
		return err
	}
	return nil
}
