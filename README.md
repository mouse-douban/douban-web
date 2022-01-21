# 正在建设中

---

# 前端部分

> 说实话，我目前没有任何前端项目开发的经验
>
> 所以我感觉这次考核对我来说是一个比较大的挑战，但也是我磨练前端技术的一个契机
>
> 毕竟我发现我平时分配给前端学习的时间属实是太少了🤣

# 后端部分

## 豆瓣 RESTful API

> 红岩寒假考核
>
> 一款仿豆瓣电影的后端 **RESTful API** 项目
>
> 前端项目地址 [Click Me!!!](https://github.com/ColdRain-Moro/RedrockWinter_Frontend)

## API 文档

+ [HTML 格式]()
+ [Markdown 格式]()
+ [OpenAPI(Swagger) 格式]()
+ [Apifox 格式]()

## 实现的Features

## 整体架构

### 分层

+ 整体分了 **4** 层：dao service controller api
+ 各层分工

> dao：和数据库IO层
>
> service:：整体的服务逻辑层
>
> controller：service 调度
>
> api：controller 调度｜入参检测

### Response

> 封装了一个Error：ServerError 和一个 RespDetail
>
> 用于返回最后的JSON response
>
> 其中 ServerError 将会作为各层中产生的 Error 来进行传递（例如dao层产生了个error就实例化一个ServerError上抛）

### 日志

> 使用标准库的 log 和 gin 自带的日志
>
> 日志会写到服务器本地文件和控制台
>
> 集群内的服务器定时会向中央日志仓库(COS)发送日志

## 分布式|集群 部署

> 一个好的代码肯定得有一个优秀的部署方案，这样才会保证服务的**高可用性**
>
> ~~更何况是我写的这种垃圾代码呢，放到实际生产环境只能靠堆配置才能勉强提供服务~~

## 学习到的东西

+ 初探分布式
