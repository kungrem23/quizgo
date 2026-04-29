package main

import (
	"log"
	"net/http"

	// "github.com/gorilla/mux"

	"github.com/kungrem23/quizgo/internal/domain/quiz"
	"github.com/kungrem23/quizgo/internal/http/handlers/auth"
	quizhandler "github.com/kungrem23/quizgo/internal/http/handlers/quiz"
	"github.com/kungrem23/quizgo/internal/platform/postgres"
)

func main() {
	db := postgres.NewDBConnection()
	defer db.Close()
	pgRepo := quiz.NewPostgresRepository(db)
	service := quiz.NewService(pgRepo)
	authHandler := auth.NewAuthHandler(service)
	quizHandler := quizhandler.NewQuizHandler(service)

	router := http.NewServeMux()
	router.HandleFunc("POST /auth/login", authHandler.Login)
	router.HandleFunc("POST /auth/register", authHandler.Register)

	router.HandleFunc("GET /api/quiz/get/{id}", quizHandler.GetQuiz)

	log.Println("server started on port :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
