package redis_cli

import (
	"time"

	"github.com/go-redis/redis"
)

var dbClient *redis.Client

func Init(hostAndPort string, poolSize int) bool {
	dbClient = redis.NewClient(&redis.Options{
		Addr:        hostAndPort,
		MaxRetries:  3,
		ReadTimeout: time.Millisecond * 1000,
		PoolSize:    poolSize,
		PoolTimeout: time.Millisecond * 300,
	})

	return true
}

func RPush(name, value string) {
	dbClient.RPush(name, value)
}

func LPop(name string) (string, error) {
	return dbClient.LPop(name).Result()
}

func Set(key, value string, expiration time.Duration) {
	dbClient.Set(key, value, expiration)
}

func Get(key string) (string, error) {
	return dbClient.Get(key).Result()
}

func Del(key string) {
	dbClient.Del(key)
}

func DelKeys(keys []string) {
	dbClient.Del(keys...)
}

func HGet(key string, field string) (string, error) {
	return dbClient.HGet(key, field).Result()
}

func Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return dbClient.Scan(cursor, match, count).Result()
}

func Pipeline() redis.Pipeliner {
	return dbClient.Pipeline()
}

func NullError(err error) bool {
	return err == redis.Nil
}
