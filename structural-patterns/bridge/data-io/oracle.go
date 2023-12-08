package main

import (
    "fmt"
    "math/rand"
)

func NewOracleDataFetcher(configStr string) IDataFetcher {
    return &OracleDataFetcher{
        Config: configStr,
    }
}

type OracleDataFetcher struct {
    Config string
}

func (o *OracleDataFetcher) Fetch(sql string) []interface{} {
    fmt.Println("Fetch data from oracle source: " + o.Config)
    rows := make([]interface{}, 0)
    // 插入两个随机数组成的切片，模拟查询要返回的数据集
    rows = append(rows, rand.Perm(10), rand.Perm(10))
    return rows
}
