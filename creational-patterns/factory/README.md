# Go设计模式--工厂模式,Go！

#### 工厂方法
##### 模式结构

###### 1.产品 （Product） 
+ 将会对接口进行声明。 
+ 对于所有由创建者及其子类构建的对象， 这些接口都是通用的。 
###### 2.具体产品 （Concrete Products） 
+ 是产品接口的不同实现。 
###### 3.创建者 （Creator） 
+ 声明返回产品对象的工厂方法。 
+ 该方法的返回对象类型必须与产品接口相匹配。 
###### 4.具体创建者 （Concrete Creators） 
+ 会重写基础工厂方法， 使其返回不同类型的产品。

#### 抽象工厂
##### 模式结构

###### 1.抽象产品 （Abstract Product）
+ 构成系列产品的一组不同但相关的产品声明接口
###### 2.具体产品 （Concrete Product）
+ 是抽象产品的多种不同类型实现。
+ 所有变体 （维多利亚/现代） 都必须实现相应的抽象产品 （椅子/沙发）。
###### 3.抽象工厂 （Abstract Factory）接口
+ 接口声明了一组创建各种抽象产品的方法。
###### 4.具体工厂 （Concrete Factory）
+ 实现抽象工厂的构建方法。
+ 每个具体工厂都对应特定产品变体， 且仅创建此种产品变体。

参考学习内容：
1. [工厂方法模式](https://refactoringguru.cn/design-patterns/factory-method)
2. [抽象工厂模式](https://refactoringguru.cn/design-patterns/factory-method)
3. [工厂模式有三个Level，你能用Go写到第几层？](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247495992&idx=1&sn=591faf1acfbd5f5aa7f0dbc95506f85c&chksm=fa8320afcdf4a9b965a768e34dff675e754de7e0c9ad95ae3134be1fa4b0990c93c6e8b3162c&scene=178&cur_album_id=2531498848431669249#rd)