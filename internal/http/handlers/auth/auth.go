package auth

import (
	// "database/sql"
	// "context"
	"encoding/json"
	"errors"
	"fmt"

	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/kungrem23/quizgo/internal/domain/quiz"
	"github.com/kungrem23/quizgo/internal/http/middleware/respond"
	// "github.com/kungrem23/quizgo/internal/utils"
)

type AuthHandler struct {
	service *quiz.Service
}

func NewAuthHandler(service *quiz.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

var loginRegex = regexp.MustCompile(`^[a-zA-Z0-9._\-#$!]{3,40}$`)

var passwordRegex = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()\-_=+\[\]{};:'",.<>\/?\\|~]{3,40}$`)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *LoginRequest) ValidateLogin() bool {
	if !loginRegex.MatchString(r.Username) {
		return false
	}
	return true
}

func (r *LoginRequest) ValidatePassword() bool {
	if !passwordRegex.MatchString(r.Password) {
		return false
	}
	return true
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respond.WriteJSON(w, http.StatusBadRequest, respond.ErrorResponse{
			Error: "bad request",
		})
		return
	}
	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)
	if !req.ValidateLogin() {
		respond.WriteJSON(w, http.StatusBadRequest, respond.ErrorResponse{
			Error: "validation failed",
			Fields: map[string]string{
				"username": "invalid",
			},
		})
		return
	}
	if !req.ValidatePassword() {
		respond.WriteJSON(w, http.StatusBadRequest, respond.ErrorResponse{
			Error: "validation failed",
			Fields: map[string]string{
				"password": "invalid",
			},
		})
		return
	}
	// user, err := h.repo.GetUserByUsername(req.Username)
	// if err == sql.ErrNoRows {
	// 	WriteJSON(w, http.StatusUnauthorized, ErrorResponse{
	// 		Error: "unauthorized",
	// 		Fields: map[string]string{
	// 			"username": "invalid",
	// 		},
	// 	})
	// 	return
	// } else if err != nil {
	// 	WriteJSON(w, http.StatusInternalServerError, ErrorResponse{
	// 		Error: "server error",
	// 	})
	// 	return
	// }
	// hash, err := utils.HashPassword(req.Password)
	// if err != nil {
	// 	WriteJSON(w, http.StatusInternalServerError, ErrorResponse{
	// 		Error: "server error",
	// 	})
	// 	return
	// }
	// if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
	// 	WriteJSON(w, http.StatusUnauthorized, ErrorResponse{
	// 		Error: "unauthorized",
	// 		Fields: map[string]string{
	// 			"password": "invalid",
	// 		},
	// 	})
	// 	return
	// }
	token, err := h.service.Login(r.Context(), req.Username, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, quiz.ErrInvalidUsername):
			respond.WriteJSON(w, http.StatusUnauthorized, respond.ErrorResponse{
				Error: "unauthorized",
				Fields: map[string]string{
					"username": "invalid",
				},
			})
			return
		case errors.Is(err, quiz.ErrInvalidPassword):
			respond.WriteJSON(w, http.StatusUnauthorized, respond.ErrorResponse{
				Error: "unauthorized",
				Fields: map[string]string{
					"password": "invalid",
				},
			})
			return
		default:
			log.Printf("Generating JWT error: %v", err)
			respond.WriteJSON(w, http.StatusInternalServerError, respond.ErrorResponse{
				Error: "server error",
			})
			return
		}
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": fmt.Sprintf("Bearer %v", token),
	})
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)
	if !req.ValidateLogin() {
		respond.WriteJSON(w, http.StatusBadRequest, respond.ErrorResponse{
			Error: "validation failed",
			Fields: map[string]string{
				"username": "invalid",
			},
		})
		return
	}
	if !req.ValidatePassword() {
		respond.WriteJSON(w, http.StatusBadRequest, respond.ErrorResponse{
			Error: "validation failed",
			Fields: map[string]string{
				"password": "invalid",
			},
		})
		return
	}
	err = h.service.Register(r.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, quiz.ErrTakenUsername) {
			respond.WriteJSON(w, http.StatusBadRequest, respond.ErrorResponse{
				Error: "registration failed",
				Fields: map[string]string{
					"username": "already exists",
				},
			})
			return
		} else {
			respond.WriteJSON(w, http.StatusInternalServerError, respond.ErrorResponse{
				Error: "registration failed",
			})
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
}
