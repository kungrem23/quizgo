package game

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/kungrem23/quizgo/internal/store/models"
	"github.com/redis/go-redis/v9"
)

type PlayerRepo struct {
	rdb *redis.Client
}

func NewPlayerRepo(rdb *redis.Client) *PlayerRepo {
	return &PlayerRepo{rdb: rdb}
}

func (r *PlayerRepo) CreateNewPlayer(username string, gameId string) (*models.Player, error) {
	ctx := context.Background()
	exists, err := r.rdb.Exists(ctx, string(gameId)).Result()
	if err != nil {
		log.Printf("Check game(id=%v) existance error: %v\n", gameId, err)
		return nil, err
	}
	if exists == 0 {
		err = r.rdb.RPush(ctx, gameId, "").Err()
		if err != nil {
			log.Printf("Pushing new game list(id=%v) error: %v\n", gameId, err)
			return nil, err
		}
	}
	player := models.Player{
		Id:       uuid.New().String(),
		Username: username,
		Score:    0,
		GameId:   gameId,
	}
	playerJSON, err := json.Marshal(player)
	if err != nil {
		log.Printf("marshall json error: %v\n", err)
		return nil, err
	}
	err = r.rdb.RPush(ctx, gameId, playerJSON).Err()
	if err != nil {
		log.Printf("pushing player to redis error: %v", err)
		return nil, err
	}
	return &player, nil
}
