package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"kevinzie/go-commerce/pkg/utils"
	"log"
	"os"
	"strconv"
)

// RedisConnection func for connect to Redis server.
func RedisConnection() (*redis.Options, *redis.Options) {
	// Define Redis database number.
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))

	// Build Redis connection URL.
	redisConnURL, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		return nil, nil
	}

	// Set Redis options.
	rdb := &redis.Options{
		Addr:     redisConnURL,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
	}

	return rdb, nil
}

func RedisData(method string, key string, value ...any) []any {
	ctx := context.Background()
	rdb, err := RedisConnection()
	if err != nil {
		log.Printf("error boss redisnya ", err)
	}
	options := redis.NewClient(rdb)

	val, errVal := options.Get(ctx, key).Result()

	if errVal == redis.Nil && method == "set" {
		log.Printf("method set redis nil")
		b, errno := json.Marshal(value[0])
		if errno != nil {
			fmt.Println(errno)
		}
		_ = options.Set(ctx, key, b, 0)
		return nil
	} else if err != nil {
		log.Printf("error", err)
		return nil
	} else {
		var tt *json.SyntaxError

		res := json.Unmarshal([]byte(val), &value)
		if errors.As(res, &tt) {
			log.Printf("value errrr", value)
			return nil
		}
		return value
	}
}
