# simple-demo

项目架构设计及各模块方案说明详细请见答辩文档：[项目结业答辩文档](https://asbmk9lxgu.feishu.cn/docx/J2o8dEJWCo8tP4xT8r3c2FPJny1)

工程无其他依赖，直接编译运行即可

```shell
go build && ./simple-demo
```

```java
项目代码结构：

--config        系统配置类
--controller    控制层
--dao           dao层
--initUtil      初始化工具类
--middleware    中间件
--public        视频存放文件夹
--service       service层
--test          测试类
  --serviceTest 服务测试类
main.go         启动类
README.md       说明文档
```

#### 3.1.1 场景分析

- 项目背景

极简抖音是一款流媒体app软件，用户可以上传和浏览视频，点赞关注，登录注册，评论等。

- 需求场景及潜在问题

1. 上传和浏览视频：浏览视频，对服务器的上行带宽要求严格，并且由于视频需要的存储空间较大，因此单独使用ftp服务器单独管理。
2. 点赞：由于是线上高并发操作，会存在多个用户对同一视频点赞，因此使用redis的单线程解决并发问题，同时为了避免用户刷赞的情况，对ip进行限流。
3. 登录：每次登录后应该记录会话状态，避免每次打开app都需要重新登录，增加操作步骤。
4. 数据更新与存储，因为存在redis和mysql两个数据库存储不同的数据，为了避免宕机造成的数据丢失，引入第三方中间件RabbitMQ进行异步更新数据。

#### 3.1.2 技术选型

+ 后端框架：Gin v1.8.2

原因：由于团队成员大多无Go语言基础，而Gin是Go世界里最流行的Web框架，因此选择Gin。

优点：Gin使用的人相对多些，文档比较好找，封装比较优雅，API友好，具有快速灵活，容错方便等优点。

+ ORM框架：Gorm v1.24.3

优点：文档齐全，对开发者友好，支持主流数据库。

+ 数据库： MySql v1.4.5 ，Redis v8.11.5

原因：成员更熟悉mysql，redis是为了线上并发问题。

优点：主流，redis操作简单，单线程+多路IO复用解决并发问题。

+ 消息队列 ：amqp v1.0.0

原因：解决数据同步的问题，并且为了避免阻塞使用消息队列进行异步通讯。

+ 测试工具：PostMan，Jemeter

原因：操作方便，测试项专业全面。

#### 3.1.3 相关开发文档

开发文档内容如下:架构方案设计，流程图，功能实现思路，功能优化点，优化后的优缺点分析等。

案例如下：

方案设计：

问题：数据一致性，以上述方案设计，如果不考虑redis会挂掉的情况（也可以部署redis集群来降低挂掉的风险），则不需要引入第三方的中间件来实现数据一致性，redis仍然是做读写主数据库，mysql只做备份，因此所有操作都是从redis中得到数据，mysql是滞后更新即可。

1. reids缓存+mq异步推送来更新mysql（需要中间件）

2. redis缓存+定时任务定时更新mysql（点赞模块实现的方案）

优缺点分析：

*本次点赞模块采用的是第二种方式，为了保证在备份时不会阻塞主线程，通过goland的**协程**与**tiker**实现定时器，实现定时任务更新mysql中的数据，避免出现redis崩溃导致数据丢失。同时在备份数据时进行分批操作，降低**
cpu**的占用。*

优点：操作简单方便，不需要依赖第三方包即可实现，并且点赞是高频率的操作，容易出现并发问题，用redis的单线程+多路IO复用来解决并发问题。

缺点：如果redis崩溃或宕机容易丢失更新时间间隔的数据，并且后期容易出现大key和大value的情况，造成查询的效率降低。

架构分析与改进：

在原有的架构中，主要存在三点问题

- 视频投稿列表存在大量的同步调用，在进行视频信息拼装时，大量同步调用会导致缓慢，用户体验较差
- 视频和图片存放在本地服务器，耦合性较高。

改进方式如下：

- 采用go语言的协程将原有同步流程异步化，提升程序性能
- 为了实现解耦，搭建FTP服务器存储视频资源，并通过SSH的方式远程调用ffmpeg截取视频关键帧
- 将FTP与SSH的链接方式改为长连接，去掉每次调用方法所需要的初始化时间

查询优化：

- 为了提高查询效率，添加外键和索引

1. 设置video表的user_id为user表的外键
2. 设置user表的username为索引
3. 设置comment表的user_id和targetId为联合索引

详情如下：

数据库选项及表设计：[项目方案设计](https://asbmk9lxgu.feishu.cn/docx/OtiadyKyno3VDGxyOLVcsouGnNO)

视频模块方案设计：[视频模块功能设计方案](https://asbmk9lxgu.feishu.cn/docx/TovBdGTEvoaDcFxXvMWcbARxnHc)

点赞模块方案设计：[点赞模块方案设计](https://asbmk9lxgu.feishu.cn/docx/ZGFOd3UgGoTTWux9TYkc9wCUnSg)

用户模块方案设计：[用户与安全模块设计说明 ](https://asbmk9lxgu.feishu.cn/docx/GbBCdRtvloZVabxfdeKc8ZcZnid)

评论模块方案设计：[评论模块设计说明](https://ghqjrqsxpx.feishu.cn/docx/JxeGd9QfHo2fXFxRrSicDQerneh) 
