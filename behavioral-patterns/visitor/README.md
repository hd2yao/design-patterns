# Go设计模式--访问者，the last！！！

### 访问者模式
> 能将算法与其所用的对象隔离开来

### 模式结构

---

###### 1.访问者 （Visitor） 接口 
+ 声明了一系列以对象结构的具体元素为参数的访问者方法
###### 2.具体访问者 （Concrete Visitor）
+ 会为不同的具体元素类实现相同行为的几个不同版本
###### 3.元素 （Element） 接口
+ 声明了一个方法来 “接收” 访问者
+ 该方法必须有一个参数被声明为访问者接口类型
###### 4.具体元素 （Concrete Element）
+ 必须实现接收方法
+ 该方法的目的是根据当前元素类将其调用重定向到相应访问者的方法
###### 5.客户端 （Client）
+ 通常会作为集合或其他复杂对象的代表


参考学习内容：
1. [访问者模式](https://refactoringguru.cn/design-patterns/visitor)
2. [Go设计模式--访客模式](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497803&idx=1&sn=4cf5d81c9d14aa00be1f52d2f30dfbb5&scene=21#wechat_redirect)