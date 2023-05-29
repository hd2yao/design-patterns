# Go设计模式--原型,继续！

#### 原型模式
##### 模式结构

###### 1.原型 （Prototype） 接口
+ 对克隆方法进行声明。
+ 在绝大多数情况下， 其中只会有一个名为 clone克隆的方法。
###### 2.具体原型 （Concrete Prototype） 类
+ 实现克隆方法。
+ 有时还需处理克隆过程中的极端情况， 例如克隆关联对象和梳理递归依赖等等。
###### 3.客户端 （Client）
+ 可以复制实现了原型接口的任何对象。



参考学习内容：
1. [原型模式](https://refactoringguru.cn/design-patterns/prototype)
2. [Go学设计模式--原型模式的考查点和使用推荐](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247496332&idx=1&sn=c28d336e9c9c964ffb78af922e0edc17&chksm=fa83231bcdf4aa0d0f7d6075fc1179a44153dc0a969c029b22d016a5ea510b1341821cc09e21&scene=178&cur_album_id=2531498848431669249#rd)