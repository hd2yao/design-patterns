package main

import (
    "fmt"
    "io"
)

func NewCsvExporter(fetcher IDataFetcher) IDataExporter {
    return &CsvExporter{
        mFetcher: fetcher,
    }
}

type CsvExporter struct {
    mFetcher IDataFetcher
}

func (c *CsvExporter) Fetcher(fetcher IDataFetcher) {
    c.mFetcher = fetcher
}

func (c *CsvExporter) Export(sql string, writer io.Writer) error {
    rows := c.mFetcher.Fetch(sql)
    fmt.Printf("CsvExporter.Export, got %v rows\n", len(rows))
    for i, row := range rows {
        fmt.Printf("  行号: %d 值: %s\n", i+1, row)
    }
    return nil
}
