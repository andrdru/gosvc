package redis

import (
	"errors"
	"time"

	redigoRedis "github.com/gomodule/redigo/redis"
)

type (
	// Redis .
	Redis interface {
		Get(key string) (data interface{}, err error)
		Setex(key string, data interface{}, ttl time.Duration) (err error)
		Set(key string, data interface{}) (err error)
		Del(key string) (err error)
		Keys(pattern string) (list []string, err error)
		ExpireAt(key string, t time.Time) (err error)
		Incr(key string) (val int64, err error)
	}

	redisImpl struct {
		pool             *redigoRedis.Pool
		operationTimeout time.Duration
	}
)

var (
	// TimeoutDefault .
	TimeoutDefault = 500 * time.Millisecond
	// NewPoolMaxIdle .
	NewPoolMaxIdle = 3
	// NewPoolIdleTimeout .
	NewPoolIdleTimeout = 240 * time.Second

	// ErrKeyNotFound .
	ErrKeyNotFound = errors.New("key not found")
	// ErrValueInvalidFormat .
	ErrValueInvalidFormat = errors.New("value format invalid")
)

// NewPool init redigo pool
func NewPool(address string) *redigoRedis.Pool {
	return &redigoRedis.Pool{
		MaxIdle:     NewPoolMaxIdle,
		IdleTimeout: NewPoolIdleTimeout,
		Dial: func() (redigoRedis.Conn, error) {
			return redigoRedis.Dial("tcp", address)
		},
		TestOnBorrow: func(c redigoRedis.Conn, t time.Time) error {
			var err error
			_, err = c.Do("PING")
			return err
		},
	}
}

// NewRedis .
func NewRedis(pool *redigoRedis.Pool, operationTimeout time.Duration) Redis {
	return &redisImpl{
		pool:             pool,
		operationTimeout: operationTimeout,
	}
}

// Get get by key
func (r *redisImpl) Get(key string) (data interface{}, err error) {
	data, err = redigoRedis.DoWithTimeout(
		r.pool.Get(),
		r.operationTimeout,
		"GET", key,
	)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, ErrKeyNotFound
	}

	return data, nil
}

// Setex set with ttl
func (r *redisImpl) Setex(key string, data interface{}, ttl time.Duration) (err error) {
	_, err = redigoRedis.DoWithTimeout(
		r.pool.Get(),
		r.operationTimeout,
		"SETEX", key, ttl.Seconds(), data,
	)

	return err
}

// Set set key value
func (r *redisImpl) Set(key string, data interface{}) (err error) {
	_, err = redigoRedis.DoWithTimeout(
		r.pool.Get(),
		r.operationTimeout,
		"SET", key, data,
	)

	return err
}

// Del delete key
func (r *redisImpl) Del(key string) (err error) {
	_, err = redigoRedis.DoWithTimeout(
		r.pool.Get(),
		r.operationTimeout,
		"DEL", key,
	)

	return err
}

// Keys get keys list
func (r *redisImpl) Keys(pattern string) (list []string, err error) {
	var data interface{}
	data, err = redigoRedis.DoWithTimeout(
		r.pool.Get(),
		r.operationTimeout,
		"KEYS", pattern,
	)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, ErrKeyNotFound
	}

	var keys, ok = data.([]interface{})
	if !ok {
		return nil, ErrKeyNotFound
	}

	for _, key := range keys {
		list = append(list, string(key.([]byte)))
	}

	return list, nil
}

func (r *redisImpl) ExpireAt(key string, t time.Time) (err error) {
	_, err = redigoRedis.DoWithTimeout(
		r.pool.Get(),
		r.operationTimeout,
		"EXPIREAT", key, t.Unix(),
	)

	return err
}

// Incr increment key
func (r *redisImpl) Incr(key string) (val int64, err error) {
	var data interface{}
	data, err = redigoRedis.DoWithTimeout(
		r.pool.Get(),
		r.operationTimeout,
		"INCR", key,
	)
	if err != nil {
		return 0, err
	}

	var ok bool
	val, ok = data.(int64)
	if !ok {
		return 0, ErrValueInvalidFormat
	}

	return val, nil
}
