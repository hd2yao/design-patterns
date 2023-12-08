# Go设计模式--桥接

##### 模式结构

---

###### 1.抽象部分 （Abstraction） 
+ 提供高层控制逻辑，依赖于完成底层实际工作的实现对象
###### 2.实现部分 （Implementation）
+ 为所有具体实现声明通用接口
+ 抽象部分仅能通过在这里声明的方法与实现对象交互
###### 3.具体实现 （Concrete Implementations）
+ 包括特定于平台的代码
###### 4.精确抽象 （Refined Abstraction）
+ 提供控制逻辑的变体

##### 适用场景
> 桥接模式 **系统多维度扩展** 和 **降低臃肿度** 
1. 在抽象和具体实现之间需要增加更多灵活性的场景
2. 一个负责某块逻辑的类存在两个或多个独立变化的维度，而这些维度都需要独立进行扩展
3. 不希望使用继承，或因为多层继承导致系统类的个数剧增


参考学习内容：
1. [桥接模式](https://refactoringguru.cn/design-patterns/bridge)
2. [Go 设计模式｜桥接模式，让代码既能多维度扩展又不会臃肿](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497649&idx=1&sn=81740f11f67e84c1aa21d701eed45ad9&scene=21#wechat_redirect)