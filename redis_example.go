package golang_redis_example

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// Redis configuration
const (
	RAW_URL    = "redis://127.0.0.1:6379"
	MAX_IDLE   = 5
	MAX_ACTIVE = 10
)

// Redis struct
type Redis struct {

	// Redis dial url
	rawUrl string

	// Maximum number of idle connection
	maxIdle int

	// Maximum number of active connection
	maxActive int
}

// Use default redis configuration
func (self *Redis) UseDefaultConfiguration() {
	self.rawUrl = RAW_URL
	self.maxIdle = MAX_IDLE
	self.maxActive = MAX_ACTIVE
}

// Get redis pool
func (self *Redis) getPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     self.maxIdle,
		MaxActive:   self.maxActive,
		Wait:        true,
		IdleTimeout: 120 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(self.rawUrl)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < 10*time.Second {
				return nil
			}

			_, err := c.Do("PING")
			return err
		},
	}
}

// Store value into redis by the given key
func (self *Redis) Put(key string, value string, time int) (string, error) {
	connection := self.getPool().Get()
	defer connection.Close()

	reply, error := redis.String(connection.Do("SETEX", key, time, value))

	return reply, error
}

// Get value from redis by the given key
func (self *Redis) Get(key string) (string, error) {
	connection := self.getPool().Get()
	defer connection.Close()

	reply, error := redis.String(connection.Do("GET", key))

	return string(reply), error
}

// Delete value from redis by the given key
func (self *Redis) Forget(key string) (string, error) {
	connection := self.getPool().Get()
	defer connection.Close()

	reply, error := redis.String(connection.Do("DEL", key))

	return string(reply), error
}
