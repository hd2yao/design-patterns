# Go设计模式--生成器,Go！

#### 生成器模式
##### 模式结构

###### 1.生成器 （Builder） 接口
+ 声明在所有类型生成器中通用的产品构造步骤。
###### 2.具体生成器 （Concrete Builders）
+ 提供构造过程的不同实现。 
+ 也可以构造不遵循通用接口的产品。
###### 3.产品 （Products）
+ 最终生成的对象。
+ 由不同生成器构造的产品无需属于同一类层次结构或接口。 
###### 4.主管 （Director） 类
+ 定义调用构造步骤的顺序。
###### 5.客户端 （Client）
+ 必须将某个生成器对象与主管类关联。


参考学习内容：
1. [生成器模式](https://refactoringguru.cn/design-patterns/builder)
2. [Go开源库、大项目的公共包，是这么用建造者模式的](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247496507&idx=1&sn=29617647f7e7c54c667abb11605c38ce&chksm=fa8322accdf4abbaa3222828ceb6f0566db26aabab529f1aa8258e8344f191867e6da88ec0d7&scene=178&cur_album_id=2531498848431669249#rd)