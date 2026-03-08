package store

import "context"

type GameStore interface {
	CreateGame(ctx context.Context) (*Game, error)
	GetGame(ctx context.Context, code string) (*Game, error)
}
