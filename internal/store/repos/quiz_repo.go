package repos

import (
	"database/sql"
	// "fmt"
	store "github.com/kungrem23/quizgo/internal/store"
	"log"
)

type QuizRepo struct {
	db *sql.DB
}

func NewQuizRepo(db *sql.DB) *QuizRepo {
	return &QuizRepo{db: db}
}

func (r *QuizRepo) CreateNewQuiz(name string, authorId int) (*store.Quiz, error) {
	query := `INSERT INTO quizzes 
	(name, author_id)
	VALUES ($1, $2)
	RETURNING id, name, author_id;`
	row := r.db.QueryRow(query, name, authorId)
	quiz := store.NewQuiz()
	err := row.Scan(&quiz.Id, &quiz.Name, &quiz.AuthorId)
	if err != nil {
		log.Printf("Adding quiz error: %v\n", err)
		return nil, err
	}
	return quiz, nil
}

func (r *QuizRepo) DeleteQuiz(id int) error {
	query := "DELETE FROM quizzes WHERE id = $1;"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Deleting quiz error: %v\n", err)
		return err
	}
	return nil
}
