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

func (r *GameRepo) GetGameById(id string) (*store.Game, error) {
	query := `SELECT id, code, is_started, quiz_id FROM games WHERE id=$1`
	row := r.db.QueryRow(query, id)
	game := store.NewGame()
	err := row.Scan(&game.Id, &game.Code, &game.IsStarted, &game.QuizId)
	if err != nil {
		log.Printf("Scanning game(id=%v) error: %v", id, err)
		return nil, err
	}
	return game, nil
}

func (r *GameRepo) GetAllGames() ([]*store.Game, error) {
	query := `SELECT id, code, is_started, quiz_id FROM games`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Get games error: %v", err)
		return nil, err
	}
	var games []*store.Game
	defer rows.Close()
	for rows.Next() {
		game := store.NewGame()
		err := rows.Scan(&game.Id, &game.Code, &game.IsStarted, &game.QuizId)
		if err != nil {
			log.Printf("Scanning game error: %v", err)
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}

func (r *GameRepo) GetGamesByQuizId(quizId int) ([]*store.Game, error) {
	query := `SELECT id, code, is_started, quiz_id FROM games
	WHERE question_id=$1`
	rows, err := r.db.Query(query, quizId)
	if err != nil {
		log.Printf("Get games error: %v", err)
		return nil, err
	}
	var games []*store.Game
	defer rows.Close()
	for rows.Next() {
		game := store.NewGame()
		err := rows.Scan(&game.Id, &game.Code, &game.IsStarted, &game.QuizId)
		if err != nil {
			log.Printf("Scanning game error: %v", err)
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}
