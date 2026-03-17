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

func (r *QuizRepo) GetQuizById(id int) (*store.Quiz, error) {
	query := `SELECT id, name, author_id FROM quizzes WHERE id=$1`
	row := r.db.QueryRow(query, id)
	quiz := store.NewQuiz()
	err := row.Scan(&quiz.Id, &quiz.Name, &quiz.AuthorId)
	if err != nil {
		log.Printf("Scanning quiz(id=%v) error: %v", id, err)
		return nil, err
	}
	return quiz, nil
}

func (r *QuizRepo) GetQuizzesByAuthorId(authorId int) ([]*store.Quiz, error) {
	query := `SELECT id, name, author_id FROM quizzes WHERE author_id=$1`
	rows, err := r.db.Query(query, authorId)
	if err != nil {
		log.Printf("Get quizzes error: %v", err)
		return nil, err
	}
	defer rows.Close()
	var quizzes []*store.Quiz
	for rows.Next() {
		quiz := store.NewQuiz()
		err = rows.Scan(&quiz.Id, &quiz.Name, &quiz.AuthorId)
		if err != nil {
			log.Printf("Scanning quiz error: %v", err)
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}
	return quizzes, nil
}

func (r *QuizRepo) GetAllQuizzes() ([]*store.Quiz, error) {
	query := `SELECT id, name, author_id FROM quizzes`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Get quizzes error: %v", err)
		return nil, err
	}
	defer rows.Close()
	var quizzes []*store.Quiz
	for rows.Next() {
		quiz := store.NewQuiz()
		err = rows.Scan(&quiz.Id, &quiz.Name, &quiz.AuthorId)
		if err != nil {
			log.Printf("Scanning quiz error: %v", err)
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}
	return quizzes, nil
}
