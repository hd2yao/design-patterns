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
