package main

import (
    "fmt"
    "time"

    "github.com/gomodule/redigo/redis"
)

// RedisCache 实现适配器接口
type RedisCache struct {
    conn *redis.Pool
}

func NewRedisCache() Cache {
    cache := &RedisCache{
        conn: &redis.Pool{
            MaxIdle:     7,
            MaxActive:   30,
            IdleTimeout: 60 * time.Second,
            Dial: func() (redis.Conn, error) {
                conn, err := redis.Dial("tcp", "localhost:6379")
                if err != nil {
                    fmt.Println(err)
                    return nil, err
                }

                if _, err := conn.Do("SELECT", 0); err != nil {
                    conn.Close()
                    fmt.Println(err)
                    return nil, err
                }

                return conn, nil
            },
        },
    }
    return cache
}

// Put 缓存数据
func (rc *RedisCache) Put(key string, value interface{}) {
    if _, err := rc.conn.Get().Do("SET", key, value); err != nil {
        fmt.Println(err)
    }
}

// Get 获取缓存中指定的 key 的值
func (rc *RedisCache) Get(key string) interface{} {
    value, err := redis.String(rc.conn.Get().Do("GET", key))
    if err != nil {
        fmt.Println(err)
        return ""
    }
    return value
}

// GetAll 从缓存中获取多个 key 的值
func (rc *RedisCache) GetAll(keys []string) map[string]interface{} {
    intKeys := make([]interface{}, len(keys))
    for i, _ := range keys {
        intKeys[i] = keys[i]
    }

    c := rc.conn.Get()
    entries := make(map[string]interface{})
    values, err := redis.Strings(c.Do("MGET", intKeys...))
    if err != nil {
        fmt.Println(err)
        return entries
    }

    for i, key := range keys {
        entries[key] = values[i]
    }

    return entries
}
