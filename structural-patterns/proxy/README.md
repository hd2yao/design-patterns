# Go设计模式--代理,go！

#### 代理模式
##### 模式结构

###### 1.服务接口 （Service Interface）
+ 代理必须遵循该接口才能伪装成服务对象
###### 2.服务 （Service）
+ 提供了一些实用的业务逻辑
###### 3.代理 （proxy）
+ 包含一个指向服务对象的引用成员变量
+ 代理完成其任务 （例如延迟初始化、 记录日志、 访问控制和缓存等） 后会将请求传递给服务对象
> 通常情况下， 代理会对其服务对象的整个生命周期进行管理。
###### 4.客户端 （Client）
+ 能通过同一接口与服务或代理进行交互

参考学习内容：
1. [单例模式](https://refactoringguru.cn/design-patterns/proxy)
2. [Go学设计模式--怕把核心代码改乱，记得用代理模式](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497211&idx=1&sn=81cd8898a26d2c54d7dee6d8cb478bb0&chksm=fa83246ccdf4ad7abe6e4f30c3a92aaa734d0ab1ac5431b05712a947ad89834729922db76039&scene=178&cur_album_id=2531498848431669249#rd)