package main

import "io"

// IDataExporter 数据导出器
type IDataExporter interface {
    Fetcher(fetcher IDataFetcher)
    Export(sql string, writer io.Writer) error
}
