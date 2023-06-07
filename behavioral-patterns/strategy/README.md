# Go设计模式--策略，如约而至！

#### 策略
##### 模式结构
###### 1.上下文 （Context）
+ 维护指向具体策略的引用
+ 仅通过`策略接口`与该对象进行交流
###### 2.策略 （Strategy） 接口
+ 是所有具体策略的通用接口， 它声明了一个上下文用于执行策略的方法
###### 3.具体策略 （Concrete Strategies）
+ 实现了上下文所用算法的各种不同变体
> 当上下文需要运行算法时，它会在其已连接的策略对象上调用执行方法。 
> 上下文不清楚其所涉及的策略类型与算法的执行方式。
###### 4.客户端 （Client）
+ 创建一个特定策略对象并将其传递给上下文
+ 上下文则会提供一个设置器以便客户端在运行时替换相关联的策略

参考学习内容：
1. [Go 程序里 if else 分支太多？试着用策略模式治理一下吧](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247496925&idx=1&sn=6535ed0a8b5ad98d56c48f568c0720ae&chksm=fa83254acdf4ac5c64ea7df4ead46fca891520a52491160e1de8ea19410dd60a6d198f5fdd65&scene=178&cur_album_id=2531498848431669249#rd)
2. [策略模式](https://refactoringguru.cn/design-patterns/strategy)