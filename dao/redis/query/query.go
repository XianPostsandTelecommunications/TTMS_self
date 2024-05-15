package redis

import "github.com/go-redis/redis/v8"

var client = &Queries{}

type Queries struct {
	redis *redis.Client
}

func New(client *redis.Client) *Queries {
	return &Queries{
		redis: client,
	}
}
