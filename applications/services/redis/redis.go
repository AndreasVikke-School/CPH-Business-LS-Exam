package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	pb "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/redis/rpc"

	"github.com/go-redis/redis/v8"
)

var (
	redis_key = "attendance_code"
)

func GetRedisClient(config Configuration) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{config.Redis.Broker},
		Password: "",
	})
}

func RandomCode() int64 {
	low := 1000000
	hi := 9999999
	rand.Seed(time.Now().UnixNano())
	return int64(low + rand.Intn(hi-low))
}

func GetUniqueCode(config Configuration) int64 {
	rdb := GetRedisClient(config)
	code := RandomCode()
	for {
		exists := rdb.HExists(rdb.Context(), redis_key, strconv.FormatInt(code, 10)).Val()
		if !exists {
			break
		}
		code = RandomCode()
	}
	return code
}

func CreateAttendanceCodeInRedis(in *pb.AttendanceCodeCreate, config Configuration) (int64, int64, float64, float64, error) {
	rdb := GetRedisClient(config)
	code := GetUniqueCode(config)
	unix := time.Now().UnixNano()/1000000 + (in.MinutesToLive * 60 * 1000)
	dataAsJson := fmt.Sprintf(`{"unix": %d, "lat": %f, "long": %f}`, unix, in.Lat, in.Long)

	_, err := rdb.HSet(rdb.Context(), redis_key, strconv.FormatInt(code, 10), dataAsJson).Result()
	if err != nil {
		log.Printf("%s", err)
		return 0, 0, 0, 0, err
	}

	return code, unix, in.Lat, in.Long, nil
}

type jsonData struct {
	Unix int64   `json:"unix,omitempty"`
	Lat  float64 `json:"lat,omitempty"`
	Long float64 `json:"long,omitempty"`
}

func GetAttendanceCodeFromRedis(code int64, config Configuration) (int64, int64, float64, float64, error) {
	rdb := GetRedisClient(config)
	exists := rdb.HExists(context.Background(), redis_key, strconv.FormatInt(code, 10)).Val()
	if !exists {
		log.Printf("code not found in redis")
		return 0, 0, 0, 0, errors.New("code not found in redis")
	}

	result, err := rdb.HGet(rdb.Context(), redis_key, strconv.FormatInt(code, 10)).Result()
	if err != nil {
		log.Printf("%s", err)
		return 0, 0, 0, 0, err
	}

	var data jsonData
	json.Unmarshal([]byte(result), &data)

	// s, _ := strconv.Unquote(string(result))
	// var data jsonData
	// if err := json.Unmarshal([]byte(s), &data); err != nil {
	// 	eh.PanicOnError(err, "Couldn't convert json result to JSON data")
	// }

	return code, data.Unix, data.Lat, data.Long, nil
}
