# Go设计模式--外观！

##### 模式结构

---

###### 1.外观 （Facade） 
+ 提供了一种访问特定子系统功能的便捷方式
###### 2.附加外观 （Additional Facade）
+ 可以避免多种不相关的功能污染单一外观
+ 客户端和其他外观都可使用附加外观
###### 3.复杂子系统 （Complex Subsystem）
+ 由数十个不同对象构成
+ 子系统类不会意识到外观的存在，它们在系统内运作并且相互之间可直接进行交互
###### 4.客户端 （Client）
+ 使用外观代替对子系统对象的直接调用


参考学习内容：
1. [外观模式](https://refactoringguru.cn/design-patterns/facade)
2. [Go 设计模式｜外观模式，一个每天都在用，却被多数人在面试中忽视的模式](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497560&idx=1&sn=c3a8f8bc5e8d7748eb16d60a242369f2&scene=21#wechat_redirect)