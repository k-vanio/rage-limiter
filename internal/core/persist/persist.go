package persist

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/k-vanio/rage-limiter/internal/domain"
)

type PersistRedis struct {
	client *redis.Client
}

func NewRedis(addr, password string, db int) *PersistRedis {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &PersistRedis{client: client}
}

func (p *PersistRedis) Store(key string, time time.Time, data interface{}) error {
	dataPersist := domain.Row{Time: time, Data: data}

	jsonData, err := json.Marshal(dataPersist)
	if err != nil {
		return err
	}

	err = p.client.LPush(context.Background(), key, jsonData).Err()
	if err != nil {
		return err
	}

	return nil
}

func (p *PersistRedis) Info(key string) []domain.Row {
	rows := make([]domain.Row, 0)

	listItems, err := p.client.LRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		return rows
	}

	for _, item := range listItems {
		var row domain.Row
		err := json.Unmarshal([]byte(item), &row)
		if err != nil {
			continue
		}

		rows = append(rows, row)
	}

	return rows
}
