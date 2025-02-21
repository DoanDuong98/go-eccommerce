package initialize

import (
	"context"
	"fmt"
	"go-ecommerce-be/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result();
	if err != nil {
		global.Logger.Error("Redis connection failed", zap.Error(err))
	}

	fmt.Println("Redis connected");
	global.Rdb = rdb
}

func RedisExample() {
	err := global.Rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	value, err := global.Rdb.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("key", value)
}
