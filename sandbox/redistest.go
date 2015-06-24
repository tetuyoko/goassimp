package main

import (
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"golang.org/x/net/context"
)

// ResourceConn adapts a Redigo connection to a Vitess Resource.
type ResourceConn struct {
	redis.Conn
}

func (r ResourceConn) Close() {
	r.Conn.Close()
}

var (
	pool *pools.ResourcePool
)

func initRedisPool() {
	pool = pools.NewResourcePool(func() (pools.Resource, error) {
		c, err := redis.Dial("tcp", ":6379")
		return ResourceConn{c}, err
	}, 1, 2, time.Minute)
	defer pool.Close()
}

func main() {
	fmt.Println("hoge")
	initRedisPool()

	ctx := context.TODO()
	resource, err := pool.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Put(resource)

	client := resource.(ResourceConn)
	n, err := client.Do("INFO")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("info=%s", n)
}
