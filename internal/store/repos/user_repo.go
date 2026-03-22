package repos

import (
	"database/sql"
	// "fmt"
	"log"

	"github.com/kungrem23/quizgo/internal/store/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateNewUser(username string, passwordHash string) (*models.User, error) {
	query := `INSERT INTO users
	(username, password_hash)
	VALUES ($1, $2)
	RETURNING id, username, password_hash;`
	row := r.db.QueryRow(query, username, passwordHash)
	user := models.NewUser()
	err := row.Scan(&user.Id, &user.Username, &user.PasswordHash)
	if err != nil {
		log.Printf("Adding user error: %v\n", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) GetAllUsers() ([]*models.User, error) {
	query := `SELECT id, username, password_hash FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Selecting all users error: %v\n", err)
		return nil, err
	}
	var users []*models.User
	defer rows.Close()
	for rows.Next() {
		user := models.NewUser()
		err := rows.Scan(&user.Id, &user.Username, &user.PasswordHash)
		if err != nil {
			log.Printf("Scanning user error: %v\n", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepo) GetUserById(id int) (*models.User, error) {
	query := `SELECT id, username, password_hash FROM users
	WHERE id = $1`
	row := r.db.QueryRow(query, id)
	user := models.NewUser()
	err := row.Scan(&user.Id, &user.Username, &user.PasswordHash)
	if err != nil {
		log.Printf("Selecting user(id=%v) error: %v\n", id, err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) GetUserByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, password_hash FROM users
	WHERE username=$1`
	row := r.db.QueryRow(query, username)
	user := models.NewUser()
	err := row.Scan(&user.Id, &user.Username, &user.PasswordHash)
	if err != nil {
		log.Printf("Selecting user by username(%v) error: %v", username, err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) Deleteuser(id int) error {
	query := "DELETE FROM users WHERE id = $1;"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Deleting user error: %v\n", err)
		return err
	}
	return nil
}
