package golang_redis_example

import "testing"

func TestRedis_UseDefaultConfiguration(t *testing.T) {
	redis := new(Redis)
	redis.UseDefaultConfiguration()

	if redis.rawUrl != "redis://127.0.0.1:6379" {
		t.Error("Redis raw url should be 'redis://127.0.0.1:6379'")
	}

	if redis.maxActive != 10 {
		t.Error("Redis max action connection should be 10")
	}

	if redis.maxIdle != 5 {
		t.Error("Redis max idle connection should be 5")
	}
}

func TestPut(t *testing.T) {
	redis := new(Redis)
	redis.UseDefaultConfiguration()

	redis.Put("redis", "redis example", 10)

	reply, _ := redis.Get("redis")

	if reply != "redis example" {
		t.Error("Message must be stored in the redis")
	}
}

func TestRedis_Forget(t *testing.T) {
	redis := new(Redis)
	redis.UseDefaultConfiguration()

	redis.Put("redis", "redis example", 10)

	reply, _ := redis.Forget("redis")

	if reply != "" {
		t.Error("Message must be empty")
	}
}
