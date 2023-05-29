# Go设计模式--单例,出发！

#### 单例模式
##### 模式结构

###### 1.单例 （Singleton）类
+ 声明了一个名为 getInstance获取实例的静态方法来返回其所属类的一个相同实例。
+ 单例的构造函数必须对客户端 （Client） 代码隐藏。 调用 `获取实例方法` 必须是获取单例对象的唯一方式。


参考学习内容：
1. [单例模式](https://refactoringguru.cn/design-patterns/singleton)
2. [最简单的单例模式，Go版本的实现你写对了吗？](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247495627&idx=1&sn=9286c6ca545280d881a3457194627cd1&chksm=fa833e5ccdf4b74addb43f60ccbddad9c3dbf650ecf2b6cfdcf58f088112fb2d6a70cd4b4b49&scene=178&cur_album_id=2531498848431669249#rd)