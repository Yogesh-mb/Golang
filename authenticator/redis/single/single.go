package single

import (
	"context"
	"time"

	"config"

	goredis "github.com/go-redis/redis/v8"
)

type Redis struct {
	Client *goredis.Client
	Expiry int
}

func NewRedis(conf config.Redis) *Redis {
	rc := goredis.NewClient(&goredis.Options{
		Addr:        conf.Host + ":" + conf.Port,
		Password:    conf.Password,
		DB:          conf.DB,
		ReadTimeOut: time.Duration(conf.ReadTimeout) * time.Millisecond,
	})

	return &Redis{Client: rc, Expiry: conf.Expiry}
}

func (r *Redis) Ping(ctx context.Context) (string, error) {
	return r.Client.Ping(ctx).Result()
}
