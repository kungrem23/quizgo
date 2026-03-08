package main

import (
	// httpApi "github.com/kungrem23/quizgo/internal/http"
	// "log"
	// "net/http"
	"fmt"
	store "github.com/kungrem23/quizgo/internal/store"
	repos "github.com/kungrem23/quizgo/internal/store/repos"
)

func main() {
	// mux := http.NewServeMux()
	// httpApi.RegisterRoutes(mux)
	// log.Println("HTTP server started on :8080")
	// if err := http.ListenAndServe(":8080", mux); err != nil {
	// 	log.Fatal(err)
	// }
	db := store.NewDBConnection()
	defer db.Close()
	quizRepo := repos.NewQuizRepo(db)
	gameRepo := repos.NewGameRepo(db)
	userRepo := repos.NewUserRepo(db)
	user, err := userRepo.CreateNewUser("testUsername4", "12345678")
	if err != nil {
		return
	}
	quiz, err := quizRepo.CreateNewQuiz("test2", user.Id)
	if err != nil {
		return
	}
	game, err := gameRepo.CreateNewGame("qwerty2", quiz.Id)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", user)
	fmt.Printf("%v\n", quiz)
	fmt.Printf("%v\n", game)
}
