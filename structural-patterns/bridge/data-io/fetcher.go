package main

// IDataFetcher 数据查询器
type IDataFetcher interface {
    Fetch(sql string) []interface{}
}
