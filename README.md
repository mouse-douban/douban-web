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

## 常见漏洞防护

+ XSS 

  > 攻击思路就是往网站里注入恶意js

  > 解决思路
  >
  > 1. 对用户的输入进行正则检测，如手机号码，邮箱地址等
  > 2. 对不能正则检测的(如用户评论)，前端做好 HTML 转义
  > 3. 对一些关键词替换，如 `javascript:` ，`<script>` 替换成 `javascript`，`script`

+ CSRF

  > 攻击思路就是在别的网站诱导用户对本站进行高危操作，如更改密码

  > 解决思路
  >
  > 1. 高危操作需要短信或者邮箱验证
  > 2. 通过包含在在请求头的JWT来认证，并且保证前端在处理外链跳转时不允许把JWT写入请求头(避免泄露)
  > 3. 不使用 cookie 保存登录令牌

  > 虽然 cookie 的 samesite 属性设置成 lax 或者 strict 能够预防 CSRF ，但前提是用户的浏览器支持 samesite

+ SQL 注入

  > 攻击思路就是注入恶意SQL语句

  > 解决思路
  >
  > 1. 严格的正则检测
  > 2. 对不能正则检测的数据使用 SQL 预处理

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

> 封装了一个Error：ServerError 和一个 RespData
>
> 用于返回最后的JSON response
>

## 集群式部署

> 一个好的代码肯定得有一个优秀的部署方案，这样才会保证服务的**高可用性**
>
> ~~更何况是我写的这种垃圾代码呢，放到实际生产环境只能靠堆配置才能勉强提供服务~~

### 前端部署

> 前端的文件全部部署到对象存储中

### 日志部署

> 使用标准库的 log 和 gin 自带的日志
>
> 日志会写到服务器本地文件和控制台
>
> 集群内的服务器定时会向中央日志仓库(COS)发送日志

## 学习到的东西

+ 初探分布式|集群部署
