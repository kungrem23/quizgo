package quiz

import (
	"context"
)

type Repository interface {
	GetQuiz(ctx context.Context, id int) (Quiz, error)
	ListQuizzes(ctx context.Context) ([]Quiz, error)
	ListByAuthor(ctx context.Context, authorID int) ([]Quiz, error)
	CreateQuiz(ctx context.Context, quiz Quiz) error
}
