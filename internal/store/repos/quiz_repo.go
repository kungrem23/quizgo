package repos

import (
	"database/sql"
	// "fmt"
	"github.com/kungrem23/quizgo/internal/store/models"
	"log"
)

type QuizRepo struct {
	db *sql.DB
}

func NewQuizRepo(db *sql.DB) *QuizRepo {
	return &QuizRepo{db: db}
}

func (r *QuizRepo) CreateNewQuiz(title string, authorId int) (*models.Quiz, error) {
	query := `INSERT INTO quizzes 
	(title, author_id)
	VALUES ($1, $2)
	RETURNING id, title, author_id;`
	row := r.db.QueryRow(query, title, authorId)
	quiz := models.NewQuiz()
	err := row.Scan(&quiz.Id, &quiz.Title, &quiz.AuthorId)
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

func (r *QuizRepo) GetQuizById(id int) (*models.Quiz, error) {
	query := `SELECT id, title, author_id FROM quizzes WHERE id=$1`
	row := r.db.QueryRow(query, id)
	quiz := models.NewQuiz()
	err := row.Scan(&quiz.Id, &quiz.Title, &quiz.AuthorId)
	if err != nil {
		log.Printf("Scanning quiz(id=%v) error: %v", id, err)
		return nil, err
	}
	return quiz, nil
}

func (r *QuizRepo) GetQuizzesByAuthorId(authorId int) ([]*models.Quiz, error) {
	query := `SELECT id, title, author_id FROM quizzes WHERE author_id=$1`
	rows, err := r.db.Query(query, authorId)
	if err != nil {
		log.Printf("Get quizzes error: %v", err)
		return nil, err
	}
	defer rows.Close()
	var quizzes []*models.Quiz
	for rows.Next() {
		quiz := models.NewQuiz()
		err = rows.Scan(&quiz.Id, &quiz.Title, &quiz.AuthorId)
		if err != nil {
			log.Printf("Scanning quiz error: %v", err)
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}
	return quizzes, nil
}

func (r *QuizRepo) GetAllQuizzes() ([]*models.Quiz, error) {
	query := `SELECT id, title, author_id FROM quizzes`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Get quizzes error: %v", err)
		return nil, err
	}
	defer rows.Close()
	var quizzes []*models.Quiz
	for rows.Next() {
		quiz := models.NewQuiz()
		err = rows.Scan(&quiz.Id, &quiz.Title, &quiz.AuthorId)
		if err != nil {
			log.Printf("Scanning quiz error: %v", err)
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}
	return quizzes, nil
}
