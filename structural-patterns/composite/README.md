# Go设计模式--组合,我又回来啦！

#### 组合模式
##### 模式结构

###### 1.组件 （Component） 接口
+ 描述了树中简单项目和复杂项目所共有的操作
###### 2.叶节点 （Leaf）
+ 树的基本结构，不包含子项目
###### 3.容器 （Container）——又名 “组合 （Composite）”
+ 包含叶节点或其他容器等子项目的单位
+ 容器不知道其子项目所属的具体类， 它只通过通用的组件接口与其子项目交互
+ 装饰基类会将所有操作委派给被封装的对象
###### 4.客户端 （Client）
+ 通过组件接口与所有项目交互

参考学习内容：
1. [组合模式](https://refactoringguru.cn/design-patterns/composite)
2. [Go 设计模式｜组合，一个对数据结构算法和职场都有提升的设计模式](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497331&idx=1&sn=cb8a154b7e6b2913cbd4db89dfe27e80&chksm=fa8327e4cdf4aef21935d4c524f26975bcd38e5e0cfb4145c751e5cca6b7e36c58306d647923&scene=178&cur_album_id=2531498848431669249#rd)