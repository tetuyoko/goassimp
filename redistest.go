package main

import (
	"fmt"
	"gopkg.in/redis.v3"
)

var RedisClient *redis.RedisClient

func main() {
	InitRedisClient()
	//	ExampleNewRedisClient()
	ExampleRedisClient()
}

func InitRedisClient() {
	RedisClient := redis.NewRedisClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func ExampleNewRedisClient() {
	pong, err := RedisClient.Ping().Result()
	//fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleRedisClient() {
	err := RedisClient.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := RedisClient.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := RedisClient.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output:
	// key
	// value
	// key2
	// does
	// not
	// exists
}
