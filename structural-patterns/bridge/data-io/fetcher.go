package main

// IDataFetcher 数据查询器
// 当前已实现 MysqlDataFetcher 和 OracleDataFetcher
// 后续要扩展支持的数据库，就新增对应的 IDataFetcher 实现即可
type IDataFetcher interface {
    Fetch(sql string) []interface{}
}
