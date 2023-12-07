package main

// 适配器接口

type Cache interface {
    Put(key string, value interface{})
    Get(key string) interface{}
    GetAll(keys []string) map[string]interface{}
}
