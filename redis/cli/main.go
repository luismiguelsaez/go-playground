package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

func RandomString(n int) string {
	var letters = []int32("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]int32, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func main() {

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	for {
		var rand_key string = RandomString(12)
		var rand_val string = RandomString(24)

		err := rdb.Set(ctx, rand_key, rand_val, 0).Err()
		if err != nil {
			log.Println("Got set error: " + err.Error())
			time.Sleep(time.Second * 1)
			continue
		}
		log.Printf("Set key [%v] value [%v]\n", rand_key, rand_val)

		val, err := rdb.Get(ctx, rand_key).Result()
		if err != nil {
			log.Println("Got get error: " + err.Error())
			time.Sleep(time.Second * 1)
			continue
		}
		log.Printf("Got key [%v] value [%v]\n", rand_key, val)

		time.Sleep(time.Second * 1)
	}
}
