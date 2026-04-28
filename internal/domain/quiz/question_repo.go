package quiz

import (
	// "database/sql"
	// "github.com/kungrem23/quizgo/internal/store/models"
	"context"
	// "log"
)

// type QuestionRepo struct {
// 	db *sql.DB
// }

// func NewQuestionRepo(db *sql.DB) *QuestionRepo {
// 	return &QuestionRepo{db: db}
// }

func (r *PostgresRepository) CreateNewQuestion(ctx context.Context, textContent string, imageId string, quizId int) error {
	queryMax := `SELECT COALESCE(MAX(position), 0) + 1
	FROM questions
	WHERE quiz_id = $1;`
	row := r.db.QueryRow(queryMax, quizId)
	var position int
	err := row.Scan(&position)
	if err != nil {
		// log.Printf("Getting question's position errror: %v\n", err)
		return err
	}
	query := `INSERT INTO questions
	(text_content, position, quiz_id, image_id)
	VALUES ($1, $2, $3, $4)`
	// question := NewQuestion()
	if imageId != "" {
		_, err = r.db.ExecContext(ctx, query, textContent, position, quizId, imageId)
	} else {
		_, err = r.db.ExecContext(ctx, query, textContent, position, quizId, nil)
	}
	// err = row.Scan(&question.Id, &question.Position, &question.TextContent, &question.QuizId, &question.ImageId)
	if err != nil {
		// log.Printf("Adding question error: %v", err)
		return err
	}
	return nil
}

func (r *PostgresRepository) DeleteQuestion(ctx context.Context, id int) error {
	question, err := r.GetQuestion(ctx, id)
	if err != nil {
		return err
	}
	pos := question.Position
	query := `DELETE FROM questions WHERE id = $1;`
	_, err = r.db.ExecContext(ctx, query, id)
	if err != nil {
		// log.Printf("Deleting quesiton error: %v\n", err)
		return err
	}
	query = `UPDATE questions
	SET position = position - 1
	WHERE position > $1;`
	_, err = r.db.Exec(query, pos)
	if err != nil {
		// log.Printf("Changing question's positions error: %v", err)
		return err
	}
	return nil
}
func (r *PostgresRepository) ChangeQuestionPosition(ctx context.Context, id int, new_position int) error {
	question, err := r.GetQuestion(ctx, id)
	if err != nil {
		return err
	}
	pos := question.Position

	query := `UPDATE questions
	SET position = position + 1
	WHERE position > $1;`
	_, err = r.db.ExecContext(ctx, query, pos)
	if err != nil {
		// log.Printf("Changing question's position error: %v\n", err)
		return err
	}
	query = `UPDATE questions
	SET position = $1
	WHERE id = $2
	RETURNING id, text_content, position, quiz_id, image_id`
	// question := NewQuestion()
	_, err = r.db.ExecContext(ctx, query, new_position, id)
	// err = row.Scan(&question.Id, &question.TextContent, &question.Position, &question.QuizId, &question.ImageId)
	if err != nil {
		// log.Printf("Changing question's(id = %v) position error: %v", id, err)
		return err
	}
	return nil
}

func (r *PostgresRepository) GetQuestion(ctx context.Context, id int) (Question, error) {
	query := `SELECT id, text_content, position, image_id, quiz_id FROM questions
	WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	var question Question
	err := row.Scan(&question.Id, &question.TextContent, &question.Position, &question.ImageId, &question.QuizId)
	// if err != nil {
	// 	log.Printf("Scanning question(id=%v) error: %v", id, err)
	// 	return question, err
	// }
	return question, err
}

func (r *PostgresRepository) GetAllQuestions(ctx context.Context) ([]Question, error) {
	query := `SELECT id, text_content, position, image_id, quiz_id FROM questions`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		// log.Printf("Get questions error: %v", err)
		return nil, err
	}
	var questions []Question
	defer rows.Close()
	for rows.Next() {
		var q Question
		err := rows.Scan(&q.Id, &q.TextContent, &q.Position, &q.ImageId, &q.QuizId)
		if err != nil {
			// log.Printf("Scanning question error: %v", err)
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

// func (r *Repository) GetQuestionsByQuizId(quiz_id int) ([]*Question, error) {
// 	query := `SELECT id, text_content, position, image_id, quiz_id FROM questions
// 	WHERE quiz_id=$1`
// 	rows, err := r.db.Query(query, quiz_id)
// 	if err != nil {
// 		log.Printf("Get questions error: %v", err)
// 		return nil, err
// 	}
// 	var questions []*Question
// 	defer rows.Close()
// 	for rows.Next() {
// 		question := NewQuestion()
// 		err := rows.Scan(&question.Id, &question.TextContent, &question.Position, &question.ImageId, &question.QuizId)
// 		if err != nil {
// 			log.Printf("Scanning question error: %v", err)
// 			return nil, err
// 		}
// 		questions = append(questions, question)
// 	}
// 	return questions, nil
// }
