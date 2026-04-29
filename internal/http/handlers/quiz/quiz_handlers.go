package quizhandler

import (
	"database/sql"
	// "encoding/json"
	"errors"
	"net/http"
	"strconv"

	// "github.com/gorilla/mux"
	"github.com/kungrem23/quizgo/internal/domain/quiz"
	"github.com/kungrem23/quizgo/internal/http/respond"
)

type QuizHandler struct {
	service *quiz.Service
}

func NewQuizHandler(service *quiz.Service) *QuizHandler {
	return &QuizHandler{service: service}
}

func (h *QuizHandler) GetQuiz(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	quiz, err := h.service.GetQuiz(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respond.WriteJSON(w, http.StatusNotFound, respond.ErrorResponse{
				Error: "not found",
			})
			return
		} else {
			respond.WriteJSON(w, http.StatusInternalServerError, respond.ErrorResponse{
				Error: "server error",
			})
			return
		}

	}
	respond.WriteJSON(w, http.StatusOK, quiz)
}
