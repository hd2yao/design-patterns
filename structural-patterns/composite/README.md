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
2. [Go学设计模式--装饰器和职责链，哪个模式实现中间件更科学？](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497282&idx=1&sn=0de76856e8649967bd3979cb122383fd&chksm=fa8327d5cdf4aec375fcae915f4eba2960b8f07f91b445f137319c81ea5a9ac4349f44d204d4&scene=178&cur_album_id=2531498848431669249#rd)