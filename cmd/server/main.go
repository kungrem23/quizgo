package main

import (
	// httpApi "github.com/kungrem23/quizgo/internal/http"
	// "log"
	// "net/http"
	"fmt"

	// "github.com/google/uuid"
	"github.com/kungrem23/quizgo/internal/store/models"
	"github.com/kungrem23/quizgo/internal/store/postgres"
	redis "github.com/kungrem23/quizgo/internal/store/redis"
	"github.com/kungrem23/quizgo/internal/store/repos"
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
	rdb, _ := redis.NewRedisConnection()
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
	err := postgres.GetTestPGData(quizRepo, quiz, userRepo, user, gameRepo, game, answerRepo, answer, questionRepo, question)
	playerRepo := repos.NewPlayerRepo(rdb)
	player, err := playerRepo.CreateNewPlayer("player2", game.Id)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", player)
}
