# Go设计模式--模板方法，我回来啦！

#### 模板方法
##### 模式结构
###### 1.抽象类 （AbstractClass）
+ 声明作为算法步骤的方法， 以及依次调用它们的实际模板方法。
+ 算法步骤可以被声明为`抽象类型`， 也可以提供一些默认实现。
###### 2.具体类 （ConcreteClass）
+ 可以重写所有步骤， 但不能重写模板方法自身。

参考学习内容：
1. [用Go学设计模式-提炼流程，减少重复开发就靠它了!](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247496813&idx=1&sn=fda31b59530deb7f2c05cff20bdc35c2&chksm=fa8325facdf4aceca2c7e4f0e38b416ea74f9b16231c06724b919ef6823de306fd3cf0c4303a&scene=178&cur_album_id=2531498848431669249#rd)
2. [模板方法模式](https://refactoringguru.cn/design-patterns/template-method)