package main

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	eh "redis_service/errorhandler"

	"github.com/go-redis/redis"
)

var (
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-cluster.redis:6379",
		Password: "",
		DB:       0,
	})
	redis_key = "attendance_code"
)

func RandomCode() int64 {
	low := 1000000
	hi := 9999999
	return int64(low + rand.Intn(hi-low))
}

func GetUniqueCode() int64 {
	code := RandomCode()
	for {
		exists := rdb.HExists(redis_key, strconv.FormatInt(code, 10)).Val()
		if exists {
			break
		}
		code = RandomCode()
	}
	return code
}

func CreateAttendanceCodeInRedis(minsToLive int64) (int64, int64, error) {
	code := GetUniqueCode()
	unix := time.Now().Unix() + (minsToLive * 60 * 1000)

	result := rdb.HSet(redis_key, strconv.FormatInt(code, 10), unix).Val()
	if !result {
		return 0, 0, errors.New("error when adding code to redis")
	}

	return code, unix, nil
}

func GetAttendanceCodeFromRedis(code int64) (int64, int64, error) {
	exists := rdb.HExists(redis_key, strconv.FormatInt(code, 10)).Val()
	if !exists {
		return 0, 0, errors.New("code not found in redis")
	}

	result := rdb.HGet(redis_key, strconv.FormatInt(code, 10)).Val()
	unix, err := strconv.ParseInt(result, 10, 64)
	eh.PanicOnError(err, "Error converting unix to int64")

	return code, unix, nil
}
