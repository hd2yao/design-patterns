# Go设计模式--责任链，AGAIN！

#### 责任链模式

##### 模式结构
###### 1.处理者 （Handler）
+ 声明了所有具体处理者的通用接口
+ 该接口通常仅包含单个方法用于请求处理， 但有时其还会包含一个设置链上下个处理者的方法
###### 2.基础处理者 （Base Handler）
+ `可选` 你可以将所有处理者共用的样本代码放置在其中
###### 3.具体处理者 （Concrete Handlers）
+ 包含处理请求的实际代码
> 每个处理者接收到请求后， 都必须决定是否进行处理， 以及是否沿着链传递请求
###### 4.客户端 （Client）
+ 可根据程序逻辑一次性或者动态地生成链
> 请求可发送给链上的任意一个处理者， 而非必须是第一个处理者。

参考学习内容：
1. [Go 程序里 if else 分支太多？试着用策略模式治理一下吧](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247496925&idx=1&sn=6535ed0a8b5ad98d56c48f568c0720ae&chksm=fa83254acdf4ac5c64ea7df4ead46fca891520a52491160e1de8ea19410dd60a6d198f5fdd65&scene=178&cur_album_id=2531498848431669249#rd)
2. [策略模式](https://refactoringguru.cn/design-patterns/strategy)