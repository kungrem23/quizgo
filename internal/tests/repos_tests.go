package tests

import (
	"log"

	"github.com/kungrem23/quizgo/internal/store/models"
	"github.com/kungrem23/quizgo/internal/store/repos"
)

func GetAllAnswersTest(r repos.AnswerRepo) ([]*models.Answer, error) {
	arr, err := r.GetAllAnswers()
	if err != nil {
		log.Printf("Get all answers error: %v", err)
		return nil, err
	}
	return arr, nil
}

func GetAnswerByIdTest(r repos.AnswerRepo, id int) (*models.Answer, error) {
	obj, err := r.GetAnswerById(id)
	if err != nil {
		log.Printf("Get answer by id error: %v", err)
		return nil, err
	}
	return obj, nil
}

func GetAnswerByQuestionIdTest(r repos.AnswerRepo, questionId int) ([]*models.Answer, error) {
	arr, err := r.GetAnswersByQuestionId(questionId)
	if err != nil {
		log.Printf("Get answers by question id error: %v", err)
		return nil, err
	}
	return arr, nil
}

func GetAllQuestionsTest(r repos.QuestionRepo) ([]*models.Question, error) {
	arr, err := r.GetAllQuestions()
	if err != nil {
		log.Printf("Get all questions test error: %v", err)
		return nil, err
	}
	return arr, nil
}

func GetQuestionByIdTest(r repos.QuestionRepo, id int) (*models.Question, error) {
	obj, err := r.GetQuestionById(id)
	if err != nil {
		log.Printf("Get question by id test error: %v", err)
		return nil, err
	}
	return obj, nil
}

func GetQuestionsByQuizIdTest(r repos.QuestionRepo, quizId int) ([]*models.Question, error) {
	arr, err := r.GetQuestionsByQuizId(quizId)
	if err != nil {
		log.Printf("Get question by quiz id test error: %v", err)
		return nil, err
	}
	return arr, nil
}

func GetAllUsersTest(r repos.UserRepo) ([]*models.User, error) {
	arr, err := r.GetAllUsers()
	if err != nil {
		log.Printf("Get all users test error: %v", err)
		return nil, err
	}
	return arr, nil
}

func GetUserByIdTest(r repos.UserRepo, id int) (*models.User, error) {
	obj, err := r.GetUserById(id)
	if err != nil {
		log.Printf("Get user by id test error: %v", err)
		return nil, err
	}
	return obj, nil
}

func GetAllGamesTest(r repos.GameRepo) ([]*models.Game, error) {
	arr, err := r.GetAllGames()
	if err != nil {
		log.Printf("Get all games test error: %v", err)
		return nil, err
	}
	return arr, nil
}

func GetGameByIdTest(r repos.GameRepo, id string) (*models.Game, error) {
	obj, err := r.GetGameById(id)
	if err != nil {
		log.Printf("Get game by id test error: %v", err)
		return nil, err
	}
	return obj, nil
}

func GetGamesByQuizIdTest(r repos.GameRepo, quizId int) ([]*models.Game, error) {
	arr, err := r.GetGamesByQuizId(quizId)
	if err != nil {
		log.Printf("Get quiz by quiz id test error: %v", err)
		return nil, err
	}
	return arr, nil
}

func GetAllQuizzesTest(r repos.QuizRepo) ([]*models.Quiz, error) {
	arr, err := r.GetAllQuizzes()
	if err != nil {
		log.Printf("Get all quizzes test error: %v", err)
		return nil, err
	}
	return arr, nil
}

func GetQuizByIdTest(r repos.QuizRepo, id int) (*models.Quiz, error) {
	obj, err := r.GetQuizById(id)
	if err != nil {
		log.Printf("Get quiz by id test error: %v", err)
		return nil, err
	}
	return obj, nil
}

func GetQuizzesByAuthorId(r repos.QuizRepo, authorId int) ([]*models.Quiz, error) {
	arr, err := r.GetQuizzesByAuthorId(authorId)
	if err != nil {
		log.Printf("Get quiz by author id test error: %v", err)
		return nil, err
	}
	return arr, nil
}
