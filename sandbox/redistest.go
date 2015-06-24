package main

import (
	"fmt"
	"gopkg.in/redis.v3"
)

var (
	RedisClient *redis.Client
)

func main() {
	Init()
	Ping()
	SetGet()
}

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Ping() {
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
}

func SetGet() {
	err := RedisClient.Set("gokey", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := RedisClient.Get("gokey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := RedisClient.Get("gokey2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("gokey2", val2)
	}
	// Output:
	// key
	// value
	// key2
	// does
	// not
	// exists
}
