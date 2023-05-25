# Go设计模式--工厂模式,Go！

#### 工厂方法
##### 模式结构

###### 1.产品 （Product） 
+ 将会对接口进行声明。 
+ 对于所有由创建者及其子类构建的对象， 这些接口都是通用的。 
###### 具体产品 （Concrete Products） 
+ 是产品接口的不同实现。 
###### 创建者 （Creator） 
+ 声明返回产品对象的工厂方法。 
+ 该方法的返回对象类型必须与产品接口相匹配。 
###### 具体创建者 （Concrete Creators） 
+ 会重写基础工厂方法， 使其返回不同类型的产品。

参考学习内容：
1. https://refactoringguru.cn/design-patterns/factory-method