# Go设计模式--享元！

#### 享元模式
##### 模式结构

###### 1.享元模式
+ 享元模式只是一种优化，在应用该模式之前，你要确定程序中存在与大量类似对象同时占用内存相关的内存问题，并且确保该问题无法使用其他更好的方式来解决。
###### 2.享元 （Flyweight）
+ 包含原始对象中部分能在多个对象中共享的状态
+ 同一享元对象可在许多不同情景中使用
+ 享元中存储的状态被称为 “内在状态”，传递给享元方法的状态被称为 “外在状态”
###### 3.情景 （Context）
+ 包含原始对象中各不相同的外在状态
+ 情景与享元对象组合在一起就能表示原始对象的全部状态

> 通常情况下， 原始对象的行为会保留在享元类中
> 
> 因此调用享元方法必须提供部分外在状态作为参数
###### 4.客户端 （Client）
+ 负责计算或存储享元的外在状态
+ 在客户端看来，享元是一种可在运行时进行配置的模板对象，具体的配置方式为向其方法中传入一些情景数据参数
###### 5.享元工厂 （Flyweight Factory）
+ 会对已有享元的缓存池进行管理
+ 客户端无需直接创建享元，只需调用工厂并向其传递目标享元的一些内在状态即可
+ 工厂会根据参数在之前已创建的享元中进行查找，如果找到满足条件的享元就将其返回；如果没有就根据参数新建享元

参考学习内容：
1. [享元模式](https://refactoringguru.cn/design-patterns/flyweight)
2. [Go学设计模式--享元模式，节省内存的好帮手](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497613&idx=1&sn=428a4ffa977421b2b2c78f36585b7c62&scene=21#wechat_redirect)