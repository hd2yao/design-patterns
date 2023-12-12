package main

import (
    "fmt"
    "math/rand"
)

func NewMysqlDataFetcher(configStr string) IDataFetcher {
    return &MysqlDataFetcher{
        Config: configStr,
    }
}

type MysqlDataFetcher struct {
    Config string
}

func (m *MysqlDataFetcher) Fetch(sql string) []interface{} {
    fmt.Println("Fetch data from mysql source: " + m.Config)
    rows := make([]interface{}, 0)
    // 插入两个随机数组成的切片，模拟咨询要返回的数据集
    rows = append(rows, rand.Perm(10), rand.Perm(10))
    return rows
}
