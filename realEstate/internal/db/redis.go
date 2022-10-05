package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	logg "realEstate/pkg/log"
)

var ctx = context.Background()

// initialize redis db
func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		logg.Warning(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		logg.Warning(err)
	}
	logg.Info("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		logg.Warning(err)
	} else {
		logg.Info("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
