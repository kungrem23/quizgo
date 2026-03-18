package repos

import (
	"database/sql"
	// "fmt"
	"log"

	"github.com/google/uuid"
	"github.com/kungrem23/quizgo/internal/store/models"
)

type GameRepo struct {
	db *sql.DB
}

func NewGameRepo(db *sql.DB) *GameRepo {
	return &GameRepo{db: db}
}

func (r *GameRepo) CreateNewGame(code string, quizId int) (*models.Game, error) {
	query := `INSERT INTO games
	(id, code, quiz_id)
	VALUES ($1, $2, $3)
	RETURNING id, code, is_started, quiz_id;`
	gameId := uuid.New().String()
	row := r.db.QueryRow(query, gameId, code, quizId)
	game := models.NewGame()
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

func (r *GameRepo) ChangeGameState(id string, isStarted bool) (*models.Game, error) {
	query := `UPDATE games
	SET is_started=$1
	WHERE id=$2
	RETURNING id, code, is_started, quiz_id`
	row := r.db.QueryRow(query, isStarted, id)
	game := models.NewGame()
	err := row.Scan(&game.Id, &game.Code, &game.IsStarted, &game.QuizId)
	if err != nil {
		log.Printf("Changing game(id=%v) state: %v", id, err)
		return nil, err
	}
	return game, nil
}

func (r *GameRepo) GetGameById(id string) (*models.Game, error) {
	query := `SELECT id, code, is_started, quiz_id FROM games WHERE id=$1`
	row := r.db.QueryRow(query, id)
	game := models.NewGame()
	err := row.Scan(&game.Id, &game.Code, &game.IsStarted, &game.QuizId)
	if err != nil {
		log.Printf("Scanning game(id=%v) error: %v", id, err)
		return nil, err
	}
	return game, nil
}

func (r *GameRepo) GetAllGames() ([]*models.Game, error) {
	query := `SELECT id, code, is_started, quiz_id FROM games`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Get games error: %v", err)
		return nil, err
	}
	var games []*models.Game
	defer rows.Close()
	for rows.Next() {
		game := models.NewGame()
		err := rows.Scan(&game.Id, &game.Code, &game.IsStarted, &game.QuizId)
		if err != nil {
			log.Printf("Scanning game error: %v", err)
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}

func (r *GameRepo) GetGamesByQuizId(quizId int) ([]*models.Game, error) {
	query := `SELECT id, code, is_started, quiz_id FROM games
	WHERE question_id=$1`
	rows, err := r.db.Query(query, quizId)
	if err != nil {
		log.Printf("Get games error: %v", err)
		return nil, err
	}
	var games []*models.Game
	defer rows.Close()
	for rows.Next() {
		game := models.NewGame()
		err := rows.Scan(&game.Id, &game.Code, &game.IsStarted, &game.QuizId)
		if err != nil {
			log.Printf("Scanning game error: %v", err)
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}
