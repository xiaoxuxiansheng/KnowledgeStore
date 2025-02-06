package redis

import (
	"context"
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisConf struct {
	MaxIdle            int
	IdleTimeoutSeconds int
	MaxActive          int
	Wait               bool
	Network            string
	Address            string
	Passwd             string
}

// Client Redis 客户端.
type Client struct {
	pool *redis.Pool
}

// GetClient 获取客户端.
func NewClient(conf *RedisConf) *Client {
	pool := getRedisPool(conf)
	return &Client{
		pool: pool,
	}
}

func (c *Client) GetConn(ctx context.Context) (redis.Conn, error) {
	return c.pool.GetContext(ctx)
}

func getRedisPool(conf *RedisConf) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     conf.MaxIdle,
		IdleTimeout: time.Duration(conf.IdleTimeoutSeconds) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := newRedisConn(conf)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		MaxActive: conf.MaxActive,
		Wait:      conf.Wait,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				panic(err)
			}
			return err
		},
	}
}

func newRedisConn(conf *RedisConf) (redis.Conn, error) {
	if conf.Address == "" {
		panic("Cannot get redis address from config")
	}

	conn, err := redis.DialContext(context.Background(),
		conf.Network, conf.Address, redis.DialPassword(conf.Passwd))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// SetEx 执行 Redis SET 命令，expireTime 时间单位为秒.
func (c *Client) SetEx(ctx context.Context, key, value string, expireSeconds int64) error {
	if key == "" || value == "" {
		return errors.New("redis SET key or value can't be empty")
	}

	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("SET", key, value, "EX", expireSeconds)
	return err
}

func (c *Client) SetNX(ctx context.Context, key, value string, expireSeconds int64) (interface{}, error) {
	if key == "" || value == "" {
		return -1, errors.New("redis SET keyNX or value can't be empty")
	}

	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return -1, err
	}
	defer conn.Close()

	reply, err := conn.Do("SETNX", key, value)
	if err != nil {
		return -1, err
	}

	r, _ := reply.(int64)
	if r == 1 {
		_, _ = conn.Do("EXPIRE", key, expireSeconds)
	}

	return reply, nil
}

// Eval 支持使用 lua 脚本.
func (c *Client) Eval(ctx context.Context, src string, keyCount int, keysAndArgs []interface{}) (interface{}, error) {
	args := make([]interface{}, 2+len(keysAndArgs))
	args[0] = src
	args[1] = keyCount
	copy(args[2:], keysAndArgs)

	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return -1, err
	}
	defer conn.Close()

	return conn.Do("EVAL", args...)
}

// Get 执行 Redis GET 命令.
func (c *Client) Get(ctx context.Context, key string) (string, error) {
	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

// Exists 执行Redis Exists 命令.
func (c *Client) Exists(ctx context.Context, keys ...string) (bool, error) {
	// redigo 对为 nil 或 empty 的参数报错信息很模糊，因此手动添加错误信息
	if len(keys) == 0 {
		return false, errors.New("redis Exists args can't be nil or empty")
	}

	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	args := make([]interface{}, len(keys))
	for i := range keys {
		args[i] = keys[i]
	}
	return redis.Bool(conn.Do("exists", args...))
}

// HGet 执行Redis HGet 命令.
func (c *Client) HGet(ctx context.Context, table, key string) (string, error) {
	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	reply, err := conn.Do("HGET", table, key)
	if err != nil {
		return "", err
	}

	if reply == nil {
		return "", nil
	}

	return redis.String(reply, err)
}

// HSet 执行Redis HSet 命令.
func (c *Client) HSet(ctx context.Context, table, key string, value interface{}) error {
	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("HSET", table, key, value)
	return err
}

func (c *Client) ZrangeByScore(ctx context.Context, table string, score1, score2 int64) ([]string, error) {
	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	raws, err := redis.Values(conn.Do("ZRANGEBYSCORE", table, score1, score2))
	if err != nil {
		return nil, err
	}

	var res []string
	for _, raw := range raws {
		tmp, ok := raw.([]byte)
		if !ok {
			continue
		}
		res = append(res, string(tmp))
	}
	return res, nil
}

// ZAdd 执行Redis ZAdd 命令.
func (c *Client) ZAdd(ctx context.Context, table string, score int64, value interface{}) error {
	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("ZADD", table, score, value)
	return err
}

func (c *Client) Expire(ctx context.Context, key string, expireSeconds int64) error {
	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("EXPIRE", key, expireSeconds)
	return err
}

func NewSetCommand(args ...interface{}) *Command {
	return &Command{
		Name: "SET",
		Args: args,
	}
}

func NewZAddCommand(args ...interface{}) *Command {
	return &Command{
		Name: "ZADD",
		Args: args,
	}
}

func NewSetBitCommand(args ...interface{}) *Command {
	return &Command{
		Name: "SETBIT",
		Args: args,
	}
}

func NewExpireCommand(args ...interface{}) *Command {
	return &Command{
		Name: "EXPIRE",
		Args: args,
	}
}

type Command struct {
	Name string
	Args []interface{}
}

func (c *Client) GetBit(ctx context.Context, key string, offset int32) (bool, error) {
	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	reply, err := redis.Int(conn.Do("GETBIT", key, offset))
	return reply == 1, err
}

// MGet 执行 Redis MGET 命令.
func (c *Client) MGet(ctx context.Context, keys ...string) ([]string, error) {
	// redigo 对为 nil 或 empty 的参数报错信息很模糊，因此手动添加错误信息
	if len(keys) == 0 {
		return nil, errors.New("redis MSET args can't be nil or empty")
	}

	conn, err := c.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	args := make([]interface{}, len(keys))
	for i := range keys {
		args[i] = keys[i]
	}
	return redis.Strings(conn.Do("MGET", args...))
}
