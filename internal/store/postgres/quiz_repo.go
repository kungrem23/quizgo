package postgres

import (
	"context"
	"database/sql"

	// "fmt"
	// "log"

	"github.com/kungrem23/quizgo/internal/domain/quiz"
	// "github.com/kungrem23/quizgo/internal/store/models"
)

type QuizRepo struct {
	db *sql.DB
}

func NewQuizRepo(db *sql.DB) *QuizRepo {
	return &QuizRepo{db: db}
}

func (r *QuizRepo) CreateQuiz(ctx context.Context, qui quiz.Quiz) error {
	query := `INSERT INTO quizzes 
	(title, author_id)
	VALUES ($1, $2)
	RETURNING (id, title, author_id)`
	// var q quiz.Quiz
	_, err := r.db.ExecContext(ctx, query, qui.Title, qui.AuthorId)
	// quiz := models.NewQuiz()
	// err := row.Scan(&quiz.Id, &quiz.Title, &quiz.AuthorId)
	if err != nil {
		// log.Printf("Adding quiz error: %v\n", err)
		return err
	}
	return nil
}

// func (r *QuizRepo) DeleteQuiz(id int) error {
// 	query := "DELETE FROM quizzes WHERE id = $1;"
// 	_, err := r.db.Exec(query, id)
// 	if err != nil {
// 		log.Printf("Deleting quiz error: %v\n", err)
// 		return err
// 	}
// 	return nil
// }

func (r *QuizRepo) GetAnswers(ctx context.Context, quizId int, questionMap map[int]*quiz.Question) error {
	query := `SELECT id, text_content, is_correct, question_id
	FROM answers
	WHERE question_id IN 
	(SELECT id FROM questions WHERE quiz_id=$1)`
	rows, err := r.db.QueryContext(ctx, query, quizId)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var a quiz.Answer
		err = rows.Scan(&a.Id, &a.TextContent, &a.IsCorrect, &a.QuestionId)
		if err != nil {
			return err
		}
		if question, ok := questionMap[a.QuestionId]; ok {
			question.Answers = append(question.Answers, a)
		}
	}
	return nil
}

func (r *QuizRepo) GetQuestions(ctx context.Context, quizId int) ([]quiz.Question, error) {
	query := `SELECT id, text_content, position, quiz_id, image_id 
	FROM questions 
	WHERE quiz_id=$1 
	ORDER BY position`
	rows, err := r.db.QueryContext(ctx, query, quizId)
	if err != nil {
		// log.Printf("Scanning quiz(id=%v) error: %v", quizId, err)
		return nil, err
	}
	defer rows.Close()
	var questions []quiz.Question
	questionMap := make(map[int]*quiz.Question)
	for rows.Next() {
		var qu quiz.Question
		err := rows.Scan(&qu.Id, &qu.TextContent, &qu.Position, &qu.QuizId, &qu.ImageId)
		if err != nil {
			return nil, err
		}
		questions = append(questions, qu)
		questionMap[qu.Id] = &questions[len(questions)-1]
	}
	err = r.GetAnswers(ctx, quizId, questionMap)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *QuizRepo) GetQuiz(ctx context.Context, id int) (quiz.Quiz, error) {
	var q quiz.Quiz
	query := `SELECT id, title, author_id FROM quizzes WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&q.Id, &q.Title, &q.AuthorId)
	if err != nil {
		// log.Printf("Scanning quiz(id=%v) error: %v", id, err)
		return q, err
	}
	questions, err := r.GetQuestions(ctx, id)
	if err != nil {
		return q, err
	}
	q.Questions = questions
	return q, nil
}

func (r *QuizRepo) ListQuizzes(ctx context.Context) ([]quiz.Quiz, error) {
	quizzes := []quiz.Quiz{}
	query := `SELECT id, title, author_id FROM quizzes`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var q quiz.Quiz
		err := rows.Scan(&q.Id, &q.Title, &q.AuthorId)
		if err != nil {
			return nil, err
		}
		questions, err := r.GetQuestions(ctx, q.Id)
		if err != nil {
			return nil, err
		}
		q.Questions = questions
		quizzes = append(quizzes, q)
	}
	return quizzes, nil
}

func (r *QuizRepo) ListByAuthor(ctx context.Context, authorID int) ([]quiz.Quiz, error) {
	quizzes := []quiz.Quiz{}
	query := `SELECT id, title, author_id FROM quizzes WHERE author_id = $1`
	rows, err := r.db.QueryContext(ctx, query, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var q quiz.Quiz
		err := rows.Scan(&q.Id, &q.Title, &q.AuthorId)
		if err != nil {
			return nil, err
		}
		questions, err := r.GetQuestions(ctx, q.Id)
		if err != nil {
			return nil, err
		}
		q.Questions = questions
		quizzes = append(quizzes, q)
	}
	return quizzes, nil
}

// func (r *QuizRepo) GetQuizzesByAuthorId(authorId int) ([]*models.Quiz, error) {
// 	query := `SELECT id, title, author_id FROM quizzes WHERE author_id=$1`
// 	rows, err := r.db.Query(query, authorId)
// 	if err != nil {
// 		log.Printf("Get quizzes error: %v", err)
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var quizzes []*models.Quiz
// 	for rows.Next() {
// 		quiz := models.NewQuiz()
// 		err = rows.Scan(&quiz.Id, &quiz.Title, &quiz.AuthorId)
// 		if err != nil {
// 			log.Printf("Scanning quiz error: %v", err)
// 			return nil, err
// 		}
// 		quizzes = append(quizzes, quiz)
// 	}
// 	return quizzes, nil
// }

// func (r *QuizRepo) GetAllQuizzes() ([]*models.Quiz, error) {
// 	query := `SELECT id, title, author_id FROM quizzes`
// 	rows, err := r.db.Query(query)
// 	if err != nil {
// 		log.Printf("Get quizzes error: %v", err)
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var quizzes []*models.Quiz
// 	for rows.Next() {
// 		quiz := models.NewQuiz()
// 		err = rows.Scan(&quiz.Id, &quiz.Title, &quiz.AuthorId)
// 		if err != nil {
// 			log.Printf("Scanning quiz error: %v", err)
// 			return nil, err
// 		}
// 		quizzes = append(quizzes, quiz)
// 	}
// 	return quizzes, nil
// }
