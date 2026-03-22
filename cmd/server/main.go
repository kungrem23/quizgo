package main

import (
	// httpApi "github.com/kungrem23/quizgo/internal/http"
	// "log"
	// "net/http"
	"fmt"

	"github.com/kungrem23/quizgo/internal/store/models"
	"github.com/kungrem23/quizgo/internal/store/postgres"
	redis "github.com/kungrem23/quizgo/internal/store/redis"
	"github.com/kungrem23/quizgo/internal/store/repos"
	"github.com/kungrem23/quizgo/internal/tests"
	// s3 "github.com/kungrem23/quizgo/internal/store/s3"
)

func main() {
	// mux := http.NewServeMux()
	// httpApi.RegisterRoutes(mux)
	// log.Println("HTTP server started on :8080")
	// if err := http.ListenAndServe(":8080", mux); err != nil {
	// 	log.Fatal(err)
	// }
	db := postgres.NewDBConnection()
	defer db.Close()
	rdb, err := redis.NewRedisConnection()
	if err != nil {
		return
	}
	defer rdb.Close()
	// s3Client, _ := s3.NewS3Connection()
	// fmt.Printf("%v\n", s3Client)

	quizRepo := repos.NewQuizRepo(db)
	userRepo := repos.NewUserRepo(db)
	gameRepo := repos.NewGameRepo(db)
	answerRepo := repos.NewAnswerRepo(db)
	questionRepo := repos.NewQuestionRepo(db)

	quiz := models.NewQuiz()
	user := models.NewUser()
	game := models.NewGame()
	answer := models.NewAnswer()
	question := models.NewQuestion()

	err = postgres.GetTestPGData(quizRepo, quiz, userRepo, user, gameRepo, game, answerRepo, answer, questionRepo, question)
	if err != nil {
		return
	}
	err = tests.GetFuncsTests(*quizRepo, *userRepo, *gameRepo, *answerRepo, *questionRepo)
	if err != nil {
		return
	}
	user1, err := userRepo.GetUserByUsername("1234567890")
	if err != nil {
		return
	}
	fmt.Printf("%v", user1)

	playerRepo := repos.NewPlayerRepo(rdb)
	player, err := playerRepo.CreateNewPlayer("player2", game.Id)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", player)
}
