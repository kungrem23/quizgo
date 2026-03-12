package repos

import (
	"database/sql"
	// "fmt"
	"log"

	"github.com/google/uuid"
	store "github.com/kungrem23/quizgo/internal/store"
)

type GameRepo struct {
	db *sql.DB
}

func NewGameRepo(db *sql.DB) *GameRepo {
	return &GameRepo{db: db}
}

func (r *GameRepo) CreateNewGame(code string, quizId int) (*store.Game, error) {
	query := `INSERT INTO games
	(id, code, quiz_id)
	VALUES ($1, $2, $3)
	RETURNING id, code, is_started, quiz_id;`
	gameId := uuid.New().String()
	row := r.db.QueryRow(query, gameId, code, quizId)
	game := store.NewGame()
	err := row.Scan(&game.Id, &game.Code, &game.IsStarted, &game.QuizId)
	if err != nil {
		log.Printf("Adding game error: %v\n", err)
		return nil, err
	}
	return game, nil
}

func (r *GameRepo) DeleteGame(id string) error {
	query := "DELETE FROM games WHERE id = $1;"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Deleting game error: %v", err)
		return err
	}
	return nil
}

func (r *GameRepo) ChangeGameState(id string, isStarted bool) (*store.Game, error) {
	query := `UPDATE games
	SET is_started=$1
	WHERE id=$2
	RETURNING id, code, is_started, quiz_id`
	row := r.db.QueryRow(query, isStarted, id)
	game := store.NewGame()
	err := row.Scan(&game.Id, &game.Code, &game.IsStarted, &game.QuizId)
	if err != nil {
		log.Printf("Changing game(id=%v) state: %v", id, err)
		return nil, err
	}
	return game, nil
}
