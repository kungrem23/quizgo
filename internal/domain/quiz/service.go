package quiz

import (
	"context"
)

type Service struct {
	repo PostgresRepository
}

func NewService(r PostgresRepository) *Service {
	return &Service{repo: r}
}

// ==============QUIZ===============

func (s *Service) GetQuiz(ctx context.Context, id int) (Quiz, error) {
	return s.repo.GetQuiz(ctx, id)
}

func (s *Service) CreateQuiz(ctx context.Context, quiz Quiz) error {
	return s.repo.CreateQuiz(ctx, quiz)
}

func (s *Service) ListQuizzes(ctx context.Context) ([]Quiz, error) {
	return s.repo.ListQuizzes(ctx)
}

func (s *Service) ListQuizzesByAuthor(ctx context.Context, authorID int) ([]Quiz, error) {
	return s.repo.ListQuizzesByAuthor(ctx, authorID)
}

// ==============QUESTION===============

func (s *Service) CreateQuestion(ctx context.Context, textContent, imageId string, quizId int) error {
	return s.repo.CreateNewQuestion(ctx, textContent, imageId, quizId)
}

func (s *Service) DeleteQuestion(ctx context.Context, id int) error {
	return s.repo.DeleteQuestion(ctx, id)
}

func (s *Service) ChangeQuestionPosition(ctx context.Context, id, new_position int) error {
	return s.repo.ChangeQuestionPosition(ctx, id, new_position)
}

func (s *Service) GetQuestion(ctx context.Context, id int) (Question, error) {
	return s.repo.GetQuestion(ctx, id)
}

func (s *Service) ListQuestons(ctx context.Context) ([]Question, error) {
	return s.repo.GetAllQuestions(ctx)
}

// ==============ANSWER===============

func (s *Service) CreateAnswer(ctx context.Context, textContent string, isCorrect bool, questionId int) error {
	return s.repo.CreateNewAnswer(ctx, textContent, isCorrect, questionId)
}

func (s *Service) DeleteAnswer(ctx context.Context, id int) error {
	return s.repo.DeleteAnswer(ctx, id)
}

func (s *Service) GetAnswer(ctx context.Context, id int) (Answer, error) {
	return s.repo.GetAnswer(ctx, id)
}

func (s *Service) ListAnswersByQuestionId(ctx context.Context, questionId int) ([]Answer, error) {
	return s.repo.GetAnswersByQuestionId(ctx, questionId)
}

// ==============IMAGE===============

func (s *Service) CreateImage(ctx context.Context, imageURL string) error {
	return s.repo.CreateNewImage(ctx, imageURL)
}

func (s *Service) DeleteImage(ctx context.Context, id string) error {
	return s.repo.DeleteImage(ctx, id)
}

func (s *Service) GetImage(ctx context.Context, id string) (Image, error) {
	return s.repo.GetImage(ctx, id)
}

// ==============USER===============

func (s *Service) CreateUser(ctx context.Context, username string, passwordHash string) error {
	return s.repo.CreateNewUser(ctx, username, passwordHash)
}

func (s *Service) DeleteUser(ctx context.Context, id int) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *Service) GetUser(ctx context.Context, id int) (User, error) {
	return s.repo.GetUser(ctx, id)
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (User, error) {
	return s.repo.GetUserByUsername(ctx, username)
}

func (s *Service) ListUsers(ctx context.Context) ([]User, error) {
	return s.repo.GetAllUsers(ctx)
}
