package mgnredis

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"golang.org/x/net/context"
)

var RedisDb *RedisDB

// ResourceConn adapts a Redigo connection to a Vitess Resource.
type ResourceConn struct {
	redis.Conn
}

func (r ResourceConn) Close() {
	r.Conn.Close()
}

type RedisDB struct {
	pool *pools.ResourcePool
}

type pooledConn struct {
	ResourceConn
	pool *pools.ResourcePool
}

func (wc pooledConn) Put() {
	wc.pool.Put(wc.ResourceConn)
}

func (db *RedisDB) Close() {
	db.pool.Close()
}

func InitRedis(server string, capacity int, maxcapacity int, duration time.Duration) {
	p := newPool(server, capacity, maxcapacity, duration)
	RedisDb = newRedisDB(p)
}

func (db *RedisDB) Ping() (interface{}, error) {
	pc, err := db.conn()
	if err != nil {
		panic(err)
	}
	defer pc.Put()

	reply, err := pc.Do("INFO")
	if err != nil {
		panic(err)
	}
	return reply, err
}

func (db *RedisDB) Set(key string, val string) (interface{}, error) {
	pc, err := db.conn()
	if err != nil {
		panic(err)
	}
	defer pc.Put()

	reply, err := pc.Do("SET", key, val)
	if err != nil {
		panic(err)
	}
	return reply, err
}

func (db *RedisDB) HSet(key string, field string, val string) (interface{}, error) {
	pc, err := db.conn()
	if err != nil {
		panic(err)
	}
	defer pc.Put()

	reply, err := pc.Do("HSET", key, field, val)
	if err != nil {
		panic(err)
	}
	return reply, err
}

func (db *RedisDB) HGetAll(key string) (interface{}, error) {
	pc, err := db.conn()
	if err != nil {
		panic(err)
	}
	defer pc.Put()

	reply, err := redis.Strings( pc.Do("HGETALL", key))
	if err != nil {
		panic(err)
	}
	return reply, err
}

func (db *RedisDB) Get(key string) (interface{}, error) {
	pc, err := db.conn()
	if err != nil {
		panic(err)
	}
	defer pc.Put()

	reply, err := redis.String(pc.Do("GET", key))
	if err != nil {
		panic(err)
	}
	return reply, err
}

// http://godoc.org/github.com/garyburd/redigo/redis#Pool
func newPool(server string, capacity int, maxcapacity int, duration time.Duration) *pools.ResourcePool {
	f := func() (pools.Resource, error) {
		c, err := redis.Dial("tcp", server)
		return ResourceConn{c}, err
	}
	return pools.NewResourcePool(f, capacity, maxcapacity, duration)
}

func newRedisDB(pool *pools.ResourcePool) *RedisDB {
	return &RedisDB{pool}
}

func (db *RedisDB) conn() (*pooledConn, error) {
	ctx := context.TODO()
	r, err := db.pool.Get(ctx)
	if err != nil {
		return nil, err
	}
	c := r.(ResourceConn)
	return &pooledConn{c, db.pool}, nil
}
