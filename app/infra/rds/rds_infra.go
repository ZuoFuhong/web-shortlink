package rds

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
	"web-shortlink/pkg/config"
)

// Infra Redis的基础设施接口
type Infra struct {
	rdb *redis.Client
}

// NewRedisInfra 创建一个redis基础设施
func NewRedisInfra() *Infra {
	rdsCfg := config.GlobalConfig().Redis
	opts, err := redis.ParseURL(rdsCfg.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	rdb := redis.NewClient(opts)
	return &Infra{
		rdb: rdb,
	}
}

func (s *Infra) Get(ctx context.Context, key string) (string, error) {
	return s.rdb.Get(ctx, key).Result()
}

func (s *Infra) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return s.rdb.Set(ctx, key, value, expiration).Err()
}

func (s *Infra) GetInt64(ctx context.Context, key string) (int64, error) {
	return s.rdb.Get(ctx, key).Int64()
}

func (s *Infra) Incr(ctx context.Context, key string) (int64, error) {
	return s.rdb.Incr(ctx, key).Result()
}

func (s *Infra) RunScript(ctx context.Context, script string, keys []string, args []interface{}) error {
	luaScript := redis.NewScript(script)
	return luaScript.Run(ctx, s.rdb, keys, args).Err()
}
