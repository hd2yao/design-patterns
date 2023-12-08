package main

import "io"

// IDataExporter 数据导出器
// 当前已实现 CsvExporter 和 JsonExporter
// 后续要扩展支持的导出格式，就新增对应的 IDataExporter 实现即可
type IDataExporter interface {
    Fetcher(fetcher IDataFetcher)
    Export(sql string, writer io.Writer) error
}
