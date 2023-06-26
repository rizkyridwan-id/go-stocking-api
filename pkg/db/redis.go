package db

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(host string, user string, password string, db int) (*redis.Client, error) {
	opt, err := redis.ParseURL(fmt.Sprintf("redis://%s:%s@%s/%d", user, password, host, db))
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	return client, nil
}
