package app

import (
	"os"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
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

const (
	RedisMaxCap     = 200
	RedisCapDefault = 20
)

// http://godoc.org/github.com/garyburd/redigo/redis#Pool
func newPool(server string) *pools.ResourcePool {
	f := func() (pools.Resource, error) {
		c, err := redis.Dial("tcp", server)
		return ResourceConn{c}, err
	}
	capacity, RedisMaxCap, idleTimeout := redisConnParams()
	return pools.NewResourcePool(f, capacity, RedisMaxCap, idleTimeout)
}

func InitRedisPool() {
	pool = newPool(":6379")
	defer pool.Close()

	//ctx := context.TODO()
	//resource, err := pool.Get(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer pool.Put(resource)

	//conn := resource.(ResourceConn)
	//defer conn.Close()

	//n, err := conn.Do("INFO")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("han")
	//log.Printf("info=%s", n)
}

func redisConnParams() (capacity int, maxCap int, idleTimeout time.Duration) {
	capacity = RedisCapDefault
	var err error

	capStr := os.Getenv("REDIS_CAPACITY")
	if capStr != "" {
		capacity, err = strconv.Atoi(capStr)
		if err != nil {
			panic(err)
		}
	}

	return capacity, RedisMaxCap, time.Minute
}
