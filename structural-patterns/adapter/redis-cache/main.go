package main

func main() {
    // 客户端在使用 Cache 时，是直接用 Cache 接口中定义的方法跟适配器交互
    // 由适配器再去转换成对三方依赖库 redisgo 的调用完成 Redis 操作
    var rc Cache
    rc = NewRedisCache()
    rc.Put("hd2yao", "love sy")

}
