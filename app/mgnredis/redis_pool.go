package mgnredis

import (
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

// http://godoc.org/github.com/garyburd/redigo/redis#Pool
func newPool(server string, capacity int, maxcapacity int, duration time.Duration) *pools.ResourcePool {
	f := func() (pools.Resource, error) {
		c, err := redis.Dial("tcp", server)
		return ResourceConn{c}, err
	}
	return pools.NewResourcePool(f, capacity, maxcapacity, duration)
}

