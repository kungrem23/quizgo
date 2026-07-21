package game

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func newGameRepoMock(t *testing.T) (*GameRepo, sqlmock.Sqlmock, func()) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sql mock: %v", err)
	}

	cleanup := func() {
		mock.ExpectClose()
		if err := db.Close(); err != nil {
			t.Errorf("close db: %v", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unmet sql expectations: %v", err)
		}
	}

	return NewGameRepo(db), mock, cleanup
}

func TestGameRepoCreateNewGame(t *testing.T) {
	repo, mock, cleanup := newGameRepoMock(t)
	defer cleanup()

	const (
		code   = "ROOM42"
		quizID = 7
		gameID = "game-id"
	)

	query := regexp.QuoteMeta(`INSERT INTO games
	(id, code, quiz_id)
	VALUES ($1, $2, $3)
	RETURNING id, code, is_started, quiz_id;`)

	mock.ExpectQuery(query).
		WithArgs(sqlmock.AnyArg(), code, quizID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "code", "is_started", "quiz_id"}).
			AddRow(gameID, code, false, quizID))

	game, err := repo.CreateNewGame(code, quizID)
	if err != nil {
		t.Fatalf("CreateNewGame returned error: %v", err)
	}
	if game.Id != gameID {
		t.Fatalf("game id = %q, want %q", game.Id, gameID)
	}
	if game.Code != code || game.IsStarted || game.QuizId != quizID {
		t.Fatalf("game = %+v, want code=%q isStarted=false quizID=%d", game, code, quizID)
	}
}

func TestGameRepoGetGameById(t *testing.T) {
	repo, mock, cleanup := newGameRepoMock(t)
	defer cleanup()

	const (
		gameID = "game-id"
		code   = "ROOM42"
		quizID = 7
	)

	query := regexp.QuoteMeta(`SELECT id, code, is_started, quiz_id FROM games WHERE id=$1`)

	mock.ExpectQuery(query).
		WithArgs(gameID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "code", "is_started", "quiz_id"}).
			AddRow(gameID, code, true, quizID))

	game, err := repo.GetGameById(gameID)
	if err != nil {
		t.Fatalf("GetGameById returned error: %v", err)
	}
	if game.Id != gameID || game.Code != code || !game.IsStarted || game.QuizId != quizID {
		t.Fatalf("game = %+v, want id=%q code=%q isStarted=true quizID=%d", game, gameID, code, quizID)
	}
}

func TestGameRepoGetGameByIdNotFound(t *testing.T) {
	repo, mock, cleanup := newGameRepoMock(t)
	defer cleanup()

	const gameID = "missing-game-id"

	query := regexp.QuoteMeta(`SELECT id, code, is_started, quiz_id FROM games WHERE id=$1`)

	mock.ExpectQuery(query).
		WithArgs(gameID).
		WillReturnError(sql.ErrNoRows)

	game, err := repo.GetGameById(gameID)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("GetGameById error = %v, want sql.ErrNoRows", err)
	}
	if game != nil {
		t.Fatalf("game = %+v, want nil", game)
	}
}

func TestGameRepoGetGamesByQuizId(t *testing.T) {
	repo, mock, cleanup := newGameRepoMock(t)
	defer cleanup()

	const quizID = 7

	query := regexp.QuoteMeta(`SELECT id, code, is_started, quiz_id FROM games
	WHERE quiz_id=$1`)

	mock.ExpectQuery(query).
		WithArgs(quizID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "code", "is_started", "quiz_id"}).
			AddRow("game-1", "ROOM1", false, quizID).
			AddRow("game-2", "ROOM2", true, quizID))

	games, err := repo.GetGamesByQuizId(quizID)
	if err != nil {
		t.Fatalf("GetGamesByQuizId returned error: %v", err)
	}
	if len(games) != 2 {
		t.Fatalf("len(games) = %d, want 2", len(games))
	}
	if games[0].Id != "game-1" || games[0].Code != "ROOM1" || games[0].IsStarted || games[0].QuizId != quizID {
		t.Fatalf("games[0] = %+v", games[0])
	}
	if games[1].Id != "game-2" || games[1].Code != "ROOM2" || !games[1].IsStarted || games[1].QuizId != quizID {
		t.Fatalf("games[1] = %+v", games[1])
	}
}
