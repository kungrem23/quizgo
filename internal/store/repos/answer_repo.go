package repos

import (
	"database/sql"
	store "github.com/kungrem23/quizgo/internal/store"
	"log"
)

type AnswerRepo struct {
	db *sql.DB
}

func NewAnswerRepo(db *sql.DB) *AnswerRepo {
	return &AnswerRepo{db: db}
}

func (r *AnswerRepo) CreateNewAnswer(textContent string, isCorrect bool, quizId int) (*store.Answer, error) {
	query := `INSERT INTO answers
	(text_content, is_correct, quiz_id)
	VALUES ($1, $2, $3)
	RETURNING id, text_content, is_correct, quiz_id;`
	row := r.db.QueryRow(query, textContent, isCorrect, quizId)
	answer := store.NewAnswer()
	err := row.Scan(&answer.Id, &answer.TextContent, &answer.IsCorrect, &answer.QuestionId)
	if err != nil {
		log.Printf("Adding answer error: %v\n", err)
		return nil, err
	}
	return answer, nil
}

func (r *AnswerRepo) DeleteAnswer(id int) error {
	query := `DELETE FROM answers WHERE id = $1;`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Deleting answer(id=%v) error: %v", id, err)
		return err
	}
	return nil
}

func (r *AnswerRepo) GetAnswerById(id int) (*store.Answer, error) {
	query := `SELECT id, text_content, is_correct, question_id FROM answers WHERE id=$1`
	row := r.db.QueryRow(query, id)
	answer := store.NewAnswer()
	err := row.Scan(&answer.Id, &answer.TextContent, &answer.IsCorrect, &answer.QuestionId)
	if err != nil {
		log.Printf("Get answer(id=%v) error: %v", id, err)
		return nil, err
	}
	return answer, nil
}

func (r *AnswerRepo) GetAllAnswers() ([]*store.Answer, error) {
	query := `SELECT id, text_content, is_correct, question_id FROM answers`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Get answers error: %v", err)
		return nil, err
	}
	var answers []*store.Answer
	defer rows.Close()
	for rows.Next() {
		answer := store.NewAnswer()
		err := rows.Scan(&answer.Id, &answer.TextContent, &answer.IsCorrect, &answer.QuestionId)
		if err != nil {
			log.Printf("Scanning answer error: %v", err)
			return nil, err
		}
		answers = append(answers, answer)
	}
	return answers, nil
}

func (r *AnswerRepo) GetAnswersByQuestionId(questionId int) ([]*store.Answer, error) {
	query := `SELECT id, text_content, is_correct, question_id FROM answers
	WHERE question_id=$1`
	rows, err := r.db.Query(query, questionId)
	if err != nil {
		log.Printf("Get answers error: %v", err)
		return nil, err
	}
	var answers []*store.Answer
	defer rows.Close()
	for rows.Next() {
		answer := store.NewAnswer()
		err := rows.Scan(&answer.Id, &answer.TextContent, &answer.IsCorrect, &answer.QuestionId)
		if err != nil {
			log.Printf("Scanning answer error: %v", err)
			return nil, err
		}
		answers = append(answers, answer)
	}
	return answers, nil
}
