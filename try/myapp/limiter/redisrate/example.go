package redis_rate

import (
	"context"
	"fmt"
	"time"

	//"github.com/go-redis/redis"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

func ExampleNewLimiter() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_ = rdb.FlushDB(ctx).Err()
	limiter := redis_rate.NewLimiter(rdb)
	for i := 0; i < 100; i++ {
		res, err := limiter.Allow(ctx, "project:123", redis_rate.PerSecond(1))

		if err != nil {
			panic(err)
		}
		for res.Allowed == 0 {
			time.Sleep(time.Millisecond * 500)
			res, err = limiter.Allow(ctx, "project:123", redis_rate.PerSecond(1))
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("i:", i, "allowed", res.Allowed, "remaining", res.Remaining)
		// Output: allowed 1 remaining 9
	}

}
