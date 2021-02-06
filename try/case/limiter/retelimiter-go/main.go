package main

import (
	"github.com/beefsack/go-rate"
	"log"
	"time"

	redis "github.com/go-redis/redis"
	ratelimiter "github.com/teambition/ratelimiter-go"
)

// Implements RedisClient for redis.Client
type redisClient struct {
	*redis.Client
}

func (c *redisClient) RateDel(key string) error {
	return c.Del(key).Err()
}
func (c *redisClient) RateEvalSha(sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	return c.EvalSha(sha1, keys, args...).Result()
}
func (c *redisClient) RateScriptLoad(script string) (string, error) {
	return c.ScriptLoad(script).Result()
}

func main() {
	rl := rate.New(3, time.Second)
	// use memory
	// limiter := ratelimiter.New(ratelimiter.Options{
	// 	Max:      10,
	// 	Duration: time.Minute, // limit to 1000 requests in 1 minute.
	// })

	// or use redis
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	limiter := ratelimiter.New(ratelimiter.Options{
		Max:      100000,
		Duration: time.Minute, // limit to 1000 requests in 1 minute.
		Client:   &redisClient{client},
	})
	//res, err := limiter.Get(r.URL.Path)

	for i := 0; i < 2000; i++ {
		rl.Wait()
		es, err := limiter.Get("123")
		log.Printf("i:%d,es:%v,err:%v", i, es, err)
	}
	select {}
}
