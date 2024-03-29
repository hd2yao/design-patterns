# Go设计模式--Next, 命令！

### 命令模式
##### 模式结构
###### 1.发送者 （Sender） 
+ 负责对请求进行初始化，其中必须包含一个成员变量来存储对于命令对象的引用
+ 发送者触发命令，而不向接收者直接发送请求
+ 发送者并不负责创建命令对象，通常会从客户端获取预先生成的命令
###### 2.命令 （Command） 接口
+ 通常仅声明一个执行命令的方法
###### 3.具体命令 （Concrete Command）
+ 实现各种类型的请求
+ 接受对象执行方法所需的参数可以声明为具体命令的成员变量
###### 4.接收者 （Receiver）
+ 包含部分业务逻辑
###### 5.客户端 （Client）
+ 创建并配置具体命令对象
+ 客户端必须将包括接收者实体在内的所有请求参数传递给命令的构造函数

参考学习内容：
1. [命令模式](https://refactoringguru.cn/design-patterns/command)
2. [Go 设计模式｜命令模式](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497999&idx=1&sn=31400313c7a80e4ea3c2271f54c08b68&scene=21#wechat_redirect)