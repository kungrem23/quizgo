package quiz

import (
	"context"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetQuiz(ctx context.Context, id int) (Quiz, error) {
	return s.repo.GetQuiz(ctx, id)
}

func (s *Service) CreateQuiz(ctx context.Context, quiz Quiz) error {
	return s.repo.CreateQuiz(ctx, quiz)
}

func (s *Service) ListQuizzes(ctx context.Context) ([]Quiz, error) {
	return s.repo.ListQuizzes(ctx)
}

func (s *Service) ListByAuthor(ctx context.Context, authorID int) ([]Quiz, error) {
	return s.repo.ListByAuthor(ctx, authorID)
}
