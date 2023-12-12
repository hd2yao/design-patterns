package main

import (
    "fmt"
    "io"
)

func NewJsonExporter(fetcher IDataFetcher) IDataExporter {
    return &JsonExporter{
        mFetcher: fetcher,
    }
}

type JsonExporter struct {
    mFetcher IDataFetcher
}

func (j *JsonExporter) Fetcher(fetcher IDataFetcher) {
    j.mFetcher = fetcher
}

func (j *JsonExporter) Export(sql string, writer io.Writer) error {
    rows := j.mFetcher.Fetch(sql)
    fmt.Printf("JsonExporter.Export, got %v rows\n", len(rows))
    for i, row := range rows {
        fmt.Printf("  行号: %d 值: %s\n", i+1, row)
    }
    return nil
}
