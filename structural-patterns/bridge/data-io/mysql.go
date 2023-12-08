package main

func NewMysqlDataFetcher(configStr string) IDataFetcher {
    return &MysqlDataFetcher{
        Config: configStr,
    }
}

type MysqlDataFetcher struct {
    Config string
}
