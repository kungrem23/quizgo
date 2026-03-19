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
	log.Printf("Got all answers succesfully")
	return arr, nil
}

func GetAnswerByIdTest(r repos.AnswerRepo, id int) (*models.Answer, error) {
	obj, err := r.GetAnswerById(id)
	if err != nil {
		log.Printf("Get answer by id error: %v", err)
		return nil, err
	}
	log.Printf("Got answer by id succesfully")
	return obj, nil
}

func GetAnswersByQuestionIdTest(r repos.AnswerRepo, questionId int) ([]*models.Answer, error) {
	arr, err := r.GetAnswersByQuestionId(questionId)
	if err != nil {
		log.Printf("Get answers by question id error: %v", err)
		return nil, err
	}
	log.Printf("Got answers by question id succesfully")
	return arr, nil
}

func GetAllQuestionsTest(r repos.QuestionRepo) ([]*models.Question, error) {
	arr, err := r.GetAllQuestions()
	if err != nil {
		log.Printf("Get all questions test error: %v", err)
		return nil, err
	}
	log.Printf("Got all questions succesfully")
	return arr, nil
}

func GetQuestionByIdTest(r repos.QuestionRepo, id int) (*models.Question, error) {
	obj, err := r.GetQuestionById(id)
	if err != nil {
		log.Printf("Get question by id test error: %v", err)
		return nil, err
	}
	log.Printf("Got question by id succesfully")
	return obj, nil
}

func GetQuestionsByQuizIdTest(r repos.QuestionRepo, quizId int) ([]*models.Question, error) {
	arr, err := r.GetQuestionsByQuizId(quizId)
	if err != nil {
		log.Printf("Get question by quiz id test error: %v", err)
		return nil, err
	}
	log.Printf("Got quesiton by quiz id succesfully")
	return arr, nil
}

func GetAllUsersTest(r repos.UserRepo) ([]*models.User, error) {
	arr, err := r.GetAllUsers()
	if err != nil {
		log.Printf("Get all users test error: %v", err)
		return nil, err
	}
	log.Printf("Got all users succesfully")
	return arr, nil
}

func GetUserByIdTest(r repos.UserRepo, id int) (*models.User, error) {
	obj, err := r.GetUserById(id)
	if err != nil {
		log.Printf("Get user by id test error: %v", err)
		return nil, err
	}
	log.Printf("Got user by id succesfully")
	return obj, nil
}

func GetAllGamesTest(r repos.GameRepo) ([]*models.Game, error) {
	arr, err := r.GetAllGames()
	if err != nil {
		log.Printf("Get all games test error: %v", err)
		return nil, err
	}
	log.Printf("Got all games succesfully")
	return arr, nil
}

func GetGameByIdTest(r repos.GameRepo, id string) (*models.Game, error) {
	obj, err := r.GetGameById(id)
	if err != nil {
		log.Printf("Get game by id test error: %v", err)
		return nil, err
	}
	log.Printf("Got game by id succesfully")
	return obj, nil
}

func GetGamesByQuizIdTest(r repos.GameRepo, quizId int) ([]*models.Game, error) {
	arr, err := r.GetGamesByQuizId(quizId)
	if err != nil {
		log.Printf("Get quiz by quiz id test error: %v", err)
		return nil, err
	}
	log.Printf("Got games by quiz id succesfully")
	return arr, nil
}

func GetAllQuizzesTest(r repos.QuizRepo) ([]*models.Quiz, error) {
	arr, err := r.GetAllQuizzes()
	if err != nil {
		log.Printf("Get all quizzes test error: %v", err)
		return nil, err
	}
	log.Printf("Got all quizzes succesfully")
	return arr, nil
}

func GetQuizByIdTest(r repos.QuizRepo, id int) (*models.Quiz, error) {
	obj, err := r.GetQuizById(id)
	if err != nil {
		log.Printf("Get quiz by id test error: %v", err)
		return nil, err
	}
	log.Printf("Got quiz by id succesfully")
	return obj, nil
}

func GetQuizzesByAuthorId(r repos.QuizRepo, authorId int) ([]*models.Quiz, error) {
	arr, err := r.GetQuizzesByAuthorId(authorId)
	if err != nil {
		log.Printf("Get quiz by author id test error: %v", err)
		return nil, err
	}
	log.Printf("Got quizzes by author id succesfully")
	return arr, nil
}

func AnswersGetFuncsTest(r repos.AnswerRepo) error {
	answers, err := GetAllAnswersTest(r)
	if err != nil {
		log.Printf("All answers get func test error: %v", err)
		return err
	}
	if len(answers) != 0 {
		_, err := GetAnswerByIdTest(r, answers[0].Id)
		if err != nil {
			log.Printf("Answer by id get func test error: %v", err)
			return err
		}
		_, err = GetAnswersByQuestionIdTest(r, answers[0].QuestionId)
		if err != nil {
			log.Printf("Answer by question id get func test error: %v", err)
			return err
		}
	}
	log.Printf("Answers get funcs tested succesfully")
	return nil
}

func GameGetFuncsTest(r repos.GameRepo) error {
	games, err := GetAllGamesTest(r)
	if err != nil {
		log.Printf("All games get func test error: %v", err)
		return err
	}
	if len(games) != 0 {
		_, err = GetGameByIdTest(r, games[0].Id)
		if err != nil {
			log.Printf("Game by id get func test error: %v", err)
			return err
		}
		_, err = GetGamesByQuizIdTest(r, games[0].QuizId)
		if err != nil {
			log.Printf("Game by quiz id get func test error: %v", err)
			return err
		}
	}
	log.Printf("Games get funcs tested succesfully")
	return nil
}

func QuestionGetFuncsTest(r repos.QuestionRepo) error {
	questions, err := GetAllQuestionsTest(r)
	if err != nil {
		log.Printf("All questions get func test error: %v", err)
		return err
	}
	if len(questions) != 0 {
		_, err := GetQuestionByIdTest(r, questions[0].Id)
		if err != nil {
			log.Printf("Question by id get func test error: %v", err)
			return err
		}
		_, err = GetQuestionsByQuizIdTest(r, questions[0].QuizId)
		if err != nil {
			log.Printf("Question by quiz id get func test error: %v", err)
			return err
		}
	}
	log.Printf("Question get funcs tested succesfully")
	return nil
}

func QuizGetFuncsTest(r repos.QuizRepo) error {
	quizzes, err := GetAllQuizzesTest(r)
	if err != nil {
		log.Printf("All quizzes get func test error: %v", err)
		return err
	}
	if len(quizzes) != 0 {
		_, err = GetQuizByIdTest(r, quizzes[0].Id)
		if err != nil {
			log.Printf("Quiz by id get func test error: %v", err)
			return err
		}
		_, err = GetQuizzesByAuthorId(r, quizzes[0].AuthorId)
		if err != nil {
			log.Printf("Quizzes by author id get func test error: %v", err)
			return err
		}
	}
	log.Printf("Quizzes get funcs tested succesfully")
	return nil
}

func UserGetFuncsTest(r repos.UserRepo) error {
	users, err := GetAllUsersTest(r)
	if err != nil {
		log.Printf("All users get func test error: %v", err)
		return err
	}
	if len(users) != 0 {
		_, err = GetUserByIdTest(r, users[0].Id)
		if err != nil {
			log.Printf("User by id get func test error: %v", err)
			return err
		}
	}
	log.Printf("Users get funcs tested succesfully")
	return nil
}

func GetFuncsTests(quizRepo repos.QuizRepo, userRepo repos.UserRepo, gameRepo repos.GameRepo,
	answerRepo repos.AnswerRepo, questionRepo repos.QuestionRepo) error {
	err := AnswersGetFuncsTest(answerRepo)
	if err != nil {
		return err
	}
	err = GameGetFuncsTest(gameRepo)
	if err != nil {
		return err
	}
	err = QuestionGetFuncsTest(questionRepo)
	if err != nil {
		return err
	}
	err = QuizGetFuncsTest(quizRepo)
	if err != nil {
		return err
	}
	err = UserGetFuncsTest(userRepo)
	if err != nil {
		return err
	}
	log.Printf("All get funcs tested succesfully!")
	return nil
}
