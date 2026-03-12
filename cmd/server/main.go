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
	rdb, _ := store.NewRedisConnection()
	defer rdb.Close()

	// quizRepo := repos.NewQuizRepo(db)
	// userRepo := repos.NewUserRepo(db)
	// gameRepo := repos.NewGameRepo(db)
	// user, err := userRepo.CreateNewUser("testUsername1", "12345678")
	// if err != nil {
	// 	return
	// }
	// quiz, err := quizRepo.CreateNewQuiz("test1", user.Id)
	// if err != nil {
	// 	return
	// }
	// game, err := gameRepo.CreateNewGame("qwerty1", quiz.Id)
	// if err != nil {
	// 	return
	// }
	// fmt.Printf("%v\n", user)
	// fmt.Printf("%v\n", quiz)
	// fmt.Printf("%v\n", game)
	playerRepo := repos.NewPlayerRepo(rdb)
	player, err := playerRepo.CreateNewPlayer("player2", "15f3e717-0593-4c76-bb46-1d9050b58b01")
	if err != nil {
		return
	}
	fmt.Printf("%v\n", player)
}
