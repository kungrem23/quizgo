package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	// "github.com/google/uuid"
	// "github.com/kungrem23/quizgo/internal/store/models"
	// "github.com/kungrem23/quizgo/internal/store/repos"
	_ "github.com/lib/pq"
)

const (
	pgHost = "localhost"
	port   = "5432"
	user   = "danilmitrosin"
	dbname = "quizgo"
)

func ConnectPG() *sql.DB {
	pgCfg := fmt.Sprintf("host=%s port=%s user=%s dbname=%s"+
		" sslmode=disable", pgHost, port, user, dbname)
	// log.Println("Connecting to PG...")
	db, err := sql.Open("postgres", pgCfg)
	if err != nil {
		log.Fatalf("Connecting to PG error: %v", err)
	}
	log.Println("Connected to PG succesfully")
	return db
}

func ApplyMigrations(db *sql.DB) {
	files, err := filepath.Glob("../../internal/store/migrations/*.sql")
	if err != nil {
		log.Fatalf("Migrations file error: %v", err)
	}
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Reading file error: %v", err)
		}
		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatalf("Applying migration error: %v", err)
		}
		log.Printf("Migration %v applied succesfully\n", file)
	}
	log.Println("All migrations applied succesfully")
}

func NewDBConnection() *sql.DB {
	db := ConnectPG()
	ApplyMigrations(db)
	return db
}

// func GetTestPGData(quizRepo *repos.QuizRepo, quiz *models.Quiz, userRepo *repos.UserRepo, user *models.User,
// 	gameRepo *repos.GameRepo, game *models.Game, answerRepo *repos.AnswerRepo, answer *models.Answer,
// 	questionRepo *repos.QuestionRepo, question *models.Question) error {
// 	user, err := userRepo.CreateNewUser(uuid.NewString(), "12345678")
// 	if err != nil {
// 		return err
// 	}
// 	quiz, err = quizRepo.CreateNewQuiz(uuid.NewString(), user.Id)
// 	if err != nil {
// 		return err
// 	}
// 	question, err = questionRepo.CreateNewQuestion(uuid.NewString(), "", quiz.Id)
// 	if err != nil {
// 		return err
// 	}
// 	answer, err = answerRepo.CreateNewAnswer(uuid.NewString(), true, question.Id)
// 	if err != nil {
// 		return err
// 	}
// 	game, err = gameRepo.CreateNewGame(uuid.NewString(), quiz.Id)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("%v\n", user)
// 	fmt.Printf("%v\n", quiz)
// 	fmt.Printf("%v\n", question)
// 	fmt.Printf("%v\n", answer)
// 	fmt.Printf("%v\n", game)
// 	return nil
// }

// func CheckGetFuncs(quizRepo *repos.QuizRepo, userRepo *repos.UserRepo, gameRepo *repos.GameRepo,
// 	answerRepo *repos.AnswerRepo, questionRepo *repos.QuestionRepo) error {

// }
