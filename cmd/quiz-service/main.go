package quizservice

import (
	"log"
	"net/http"

	// "github.com/gorilla/mux"

	"github.com/kungrem23/quizgo/internal/domain/quiz"
	"github.com/kungrem23/quizgo/internal/http/handlers/auth"
	quizhandler "github.com/kungrem23/quizgo/internal/http/handlers/quiz"
	"github.com/kungrem23/quizgo/internal/http/middleware"
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

	router.Handle("POST /api/quizzes", middleware.Auth(http.HandlerFunc(quizHandler.CreateQuiz)))
	router.HandleFunc("GET /api/quizzes/{id}", quizHandler.GetQuiz)
	router.HandleFunc("GET /api/quizzes/", quizHandler.ListQuizzes)
	router.HandleFunc("GET /api/users/{authorId}/quizzes", quizHandler.ListQuizzesByAuthor)

	log.Println("server started on port :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
