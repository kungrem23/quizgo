package quiz

import (
	"context"
	"database/sql"
)

type Repository interface {
	// ========QUIZ======
	GetQuiz(ctx context.Context, id int) (Quiz, error)
	ListQuizzes(ctx context.Context) ([]Quiz, error)
	ListByAuthor(ctx context.Context, authorID int) ([]Quiz, error)
	CreateQuiz(ctx context.Context, quiz Quiz) error

	// =========QUESTION=========
	CreateQuestion(ctx context.Context, textContent, imageId string, quizId int) error
	DeleteQuestion(ctx context.Context, id int) error
	ChangeQuestionPosition(ctx context.Context, id, new_position int) error
	GetQuestion(ctx context.Context, id int) (Question, error)
	ListQuestons(ctx context.Context) ([]Question, error)

	// ==========ANSWER===========

	CreateAnswer(ctx context.Context, textContent string, isCorrect bool, questionId int) error
	DeleteAnswer(ctx context.Context, id int) error
	GetAnswer(ctx context.Context, id int) (Answer, error)
	ListAnswersByQuestionId(ctx context.Context, questionId int) ([]Answer, error)

	// ==========IMAGE============

	CreateImage(ctx context.Context, imageURL string) error
	DeleteImage(ctx context.Context, id string) error
	GetImage(ctx context.Context, id string) (Image, error)

	// ==========USER=============

	CreateUser(ctx context.Context, username string, passwordHash string) error
	DeleteUser(ctx context.Context, id int) error
	GetUser(ctx context.Context, id int) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}
