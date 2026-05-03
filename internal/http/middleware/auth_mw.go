package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kungrem23/quizgo/internal/http/middleware/respond"
	"github.com/kungrem23/quizgo/internal/utils"
)

var ErrInvalidToken error = errors.New("Invalid token")
var ErrInsufficientRights error = errors.New("Insufficient rights")

type contextKey string

const userIDKey contextKey = "user_id"

func UserIDFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(userIDKey).(int)
	return id, ok
}

func userIDFromToken(token *jwt.Token) (int, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, ErrInvalidToken
	}
	idValue, ok := claims["id"]
	if !ok {
		return 0, ErrInvalidToken
	}
	idFloat, ok := idValue.(float64)
	if !ok {
		return 0, ErrInvalidToken
	}
	return int(idFloat), nil
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			respond.WriteJSON(w, http.StatusUnauthorized, respond.ErrorResponse{
				Error: "invalid token",
			})
			return
		}
		token, err := utils.ParseJWT(parts[1])
		if err != nil || !token.Valid {
			respond.WriteJSON(w, http.StatusUnauthorized, respond.ErrorResponse{
				Error: "invalid token",
			})
			return
		}
		userID, err := userIDFromToken(token)
		if err != nil {
			respond.WriteJSON(w, http.StatusUnauthorized, respond.ErrorResponse{
				Error: "invalid token",
			})
			return
		}
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
