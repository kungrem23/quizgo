package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/kungrem23/quizgo/internal/store/repos"
	"github.com/kungrem23/quizgo/internal/utils"
)

type AuthHandler struct {
	userRepo *repos.UserRepo
}

func NewAuthHandler(userRepo *repos.UserRepo) *AuthHandler {
	return &AuthHandler{userRepo: userRepo}
}

var loginRegex = regexp.MustCompile(`^[a-zA-Z0-9._-#$!]{3-40}$`)

var passwordRegex = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()\-_=+\[\]{};:'",.<>\/?\\|~]{3-40}$`)

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

type ErrorResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message,omitempty"`
	Fields  map[string]string `json:"fields,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)
	if !req.ValidateLogin() {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{
			Error: "validation failed",
			Fields: map[string]string{
				"username": "invalid",
			},
		})
		return
	}
	if !req.ValidatePassword() {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{
			Error: "validation failed",
			Fields: map[string]string{
				"password": "invalid",
			},
		})
		return
	}
	user, err := h.userRepo.GetUserByUsername(req.Username)
	if err == sql.ErrNoRows {
		WriteJSON(w, http.StatusUnauthorized, ErrorResponse{
			Error: "unauthorized",
			Fields: map[string]string{
				"username": "invalid",
			},
		})
		return
	} else if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{
			Error: "server error",
		})
		return
	}
	// hash, err := utils.HashPassword(req.Password)
	// if err != nil {
	// 	WriteJSON(w, http.StatusInternalServerError, ErrorResponse{
	// 		Error: "server error",
	// 	})
	// 	return
	// }
	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		WriteJSON(w, http.StatusUnauthorized, ErrorResponse{
			Error: "unauthorized",
			Fields: map[string]string{
				"password": "invalid",
			},
		})
		return
	}
	token, err := utils.GenerateJWT(user.Id)
	if err != nil {
		log.Printf("Generating JWT error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{
			Error: "server error",
		})
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": fmt.Sprintf("Bearer %v", token),
	})
	return
}
