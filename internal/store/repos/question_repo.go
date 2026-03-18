package repos

import (
	"database/sql"
	"github.com/kungrem23/quizgo/internal/store/models"
	"log"
)

type QuestionRepo struct {
	db *sql.DB
}

func NewQuestionRepo(db *sql.DB) *QuestionRepo {
	return &QuestionRepo{db: db}
}

func (r *QuestionRepo) CreateNewQuestion(textContent string, imageId string, quizId int) (*models.Question, error) {
	queryMax := `SELECT COALESCE(MAX(position), 0) + 1
	FROM questions
	WHERE quiz_id = $1;`
	row := r.db.QueryRow(queryMax, quizId)
	var position int
	err := row.Scan(&position)
	if err != nil {
		log.Printf("Getting question's position errror: %v\n", err)
		return nil, err
	}
	query := `INSERT INTO questions
	(text_content, position, quiz_id, image_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id, position, text_content, quiz_id, image_id;`
	question := models.NewQuestion()
	if imageId != "" {
		row = r.db.QueryRow(query, textContent, position, quizId, imageId)
	} else {
		row = r.db.QueryRow(query, textContent, position, quizId, nil)
	}
	err = row.Scan(&question.Id, &question.Position, &question.TextContent, &question.QuizId, &question.ImageId)
	if err != nil {
		log.Printf("Adding question error: %v", err)
		return nil, err
	}
	return question, nil
}

func (r *QuestionRepo) DeleteQuestion(id int) error {
	query := `DELETE FROM questions WHERE id = $1;`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Deleting quesiton error: %v\n", err)
		return err
	}
	query = `UPDATE questions
	SET position = position - 1
	WHERE position > $1;`
	_, err = r.db.Exec(query, id)
	if err != nil {
		log.Printf("Changing question's positions error: %v", err)
		return err
	}
	return nil
}
func (r *QuestionRepo) ChangeQuestionPosition(id int, new_position int) (*models.Question, error) {
	query := `UPDATE questions
	SET position = position + 1
	WHERE position > $1;`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Changing question's position error: %v\n", err)
		return nil, err
	}
	query = `UPDATE questions
	SET position = $1
	WHERE id = $2
	RETURNING id, text_content, position, quiz_id, image_id`
	question := models.NewQuestion()
	row := r.db.QueryRow(query, new_position, id)
	err = row.Scan(&question.Id, &question.TextContent, &question.Position, &question.QuizId, &question.ImageId)
	if err != nil {
		log.Printf("Changing question's(id = %v) position error: %v", id, err)
		return nil, err
	}
	return question, nil
}

func (r *QuestionRepo) GetQuestionById(id int) (*models.Question, error) {
	query := `SELECT id, text_content, position, image_id, quiz_id FROM questions
	WHERE id=$1`
	row := r.db.QueryRow(query, id)
	question := models.NewQuestion()
	err := row.Scan(&question.Id, &question.TextContent, &question.Position, &question.ImageId, &question.QuizId)
	if err != nil {
		log.Printf("Scanning question(id=%v) error: %v", id, err)
		return nil, err
	}
	return question, nil
}

func (r *QuestionRepo) GetAllQuestions() ([]*models.Question, error) {
	query := `SELECT id, text_content, position, image_id, quiz_id FROM questions`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Get questions error: %v", err)
		return nil, err
	}
	var questions []*models.Question
	defer rows.Close()
	for rows.Next() {
		question := models.NewQuestion()
		err := rows.Scan(&question.Id, &question.TextContent, &question.Position, &question.ImageId, &question.QuizId)
		if err != nil {
			log.Printf("Scanning question error: %v", err)
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}

func (r *QuestionRepo) GetQuestionsByQuizId(quiz_id int) ([]*models.Question, error) {
	query := `SELECT id, text_content, position, image_id, quiz_id FROM questions
	WHERE quiz_id=$1`
	rows, err := r.db.Query(query, quiz_id)
	if err != nil {
		log.Printf("Get questions error: %v", err)
		return nil, err
	}
	var questions []*models.Question
	defer rows.Close()
	for rows.Next() {
		question := models.NewQuestion()
		err := rows.Scan(&question.Id, &question.TextContent, &question.Position, &question.ImageId, &question.QuizId)
		if err != nil {
			log.Printf("Scanning question error: %v", err)
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}
