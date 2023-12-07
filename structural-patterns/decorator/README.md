# Go设计模式--装饰,顺便迎接小长假！

#### 装饰模式
##### 模式结构

###### 1.部件 （Component）
+ 声明封装器和被封装对象的公用接口
###### 2.具体部件 （Concrete Component）类
+ 被封装对象所属的类
+ 定义了基础行为， 但装饰类可以改变这些行为
###### 3.基础装饰 （Base Decorator）类
+ 拥有一个指向被封装对象的引用成员变量
+ 该变量的类型应当被声明为通用部件接口， 这样它就可以引用具体的部件和装饰
+ 装饰基类会将所有操作委派给被封装的对象
###### 4.具体装饰类 （Concrete Decorators）
+ 定义了可动态添加到部件的额外行为
+ 具体装饰类会重写装饰基类的方法， 并在调用父类方法之前或之后进行额外的行为
+ ###### 5.客户端 （Client）
+ 可以使用多层装饰来封装部件， 只要它能使用通用接口与所有对象互动即可

参考学习内容：
1. [装饰模式](https://refactoringguru.cn/design-patterns/decorator)
2. [Go学设计模式--装饰器和职责链，哪个模式实现中间件更科学？](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497282&idx=1&sn=0de76856e8649967bd3979cb122383fd&chksm=fa8327d5cdf4aec375fcae915f4eba2960b8f07f91b445f137319c81ea5a9ac4349f44d204d4&scene=178&cur_album_id=2531498848431669249#rd)