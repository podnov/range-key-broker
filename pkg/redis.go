package pkg

import (
	"fmt"
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"time"
)

func createMutexes(names []string) map[string]*redsync.Mutex {
	result := map[string]*redsync.Mutex{}

	r := createRedsync()
	mutexOptions := []redsync.Option{
		redsync.SetExpiry(6 * time.Hour),
		redsync.SetTries(1),
	}

	for _, name := range names {
		mutex := r.NewMutex(name, mutexOptions...)
		result[name] = mutex
	}

	return result
}

func createPool() *redis.Pool {
	return &redis.Pool{
		MaxActive:   20,
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			address := viper.GetString(ConfigKeyRedisAddress)

			result, err := redis.Dial("tcp", address)
			if err != nil {
				message := fmt.Sprintf("redis.Dial failed for [%s]", address)
				return nil, errors.Wrap(err, message)
			}

			password := viper.GetString(ConfigKeyRedisPassword)
			if password != "" {
				if _, err := result.Do("AUTH", password); err != nil {
					result.Close()
					message := fmt.Sprintf("redis.Do Auth failed")
					return nil, errors.Wrap(err, message)
				}
			}

			return result, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func createRedsync() *redsync.Redsync {
	pool := createPool()
	pools := []redsync.Pool{pool}
	return redsync.New(pools)
}
