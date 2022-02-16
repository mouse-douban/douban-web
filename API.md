# 豆瓣API

> v1.0.0

# 电影

## GET 获取电影列表

GET /subjects

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|start|query|string|false|开始的索引，不填为0|
|limit|query|string|false|限制数量，不填为20|
|sort|query|string|false|排序规则｜填 hotest|latest，不填为 latest|
|tag|query|string|false|标签 喜剧,生活,爱情,动作,科幻,悬疑,惊悚,动画,奇幻|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 电影/获取信息

## GET 获取电影的短评列表

GET /subjects/{id}/comments

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|start|query|string|false|开始序列号，不填默认为0|
|limit|query|string|false|数量限制，不填为20|
|sort|query|string|false|排序规则｜填 hotest(最热门)|latest(最新)，不填为hotest|
|type|query|string|false|类型｜想看 before 看过 after|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取电影的主页信息

GET /subjects/{id}

+ 请求中 scope 有什么，返回的 json 的 data 里面就有什么字段
+ 不建议一次性把 scope 填满
+ 没有 scope 视为空字符串
+ 其中 comments 返回的是 看过 的短评

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|电影的id|
|scope|query|string|false|范围, 填comments|reviews|discussions|

> 返回示例

> 成功

```json
{
  "status": 20000,
  "info": "success",
  "data": {
    "mid": 1,
    "name": "千与千寻",
    "stars": 5,
    "date": "2022-02-02T14:47:00Z",
    "tags": "动画,奇幻",
    "avatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2557573348.jpg",
    "detail": {
      "nicknames": [
        "千与千寻"
      ],
      "director": "宫崎骏",
      "writers": [
        "宫崎骏"
      ],
      "characters": [
        "I dont know"
      ],
      "type": [
        "动画"
      ],
      "region": "日本",
      "language": "日语",
      "release": "2022-02-03T21:23:47.013195+08:00",
      "period": 117,
      "IMDb": "IAwHqAwDAndaHsa"
    },
    "score": {
      "score": "9.7",
      "total_cnt": 900,
      "five": "85%",
      "four": "10%",
      "three": "3%",
      "two": "1%",
      "one": "1%"
    },
    "plot": "千与千寻是一个好看的电影"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» status|integer|true|none|none|
|» info|string|true|none|none|
|» data|object|true|none|none|
|»» mid|integer|true|none|none|
|»» name|string|true|none|none|
|»» stars|integer|true|none|none|
|»» date|string|true|none|none|
|»» tags|string|true|none|none|
|»» avatar|string|true|none|none|
|»» detail|object|true|none|none|
|»»» nicknames|[string]|true|none|none|
|»»» director|string|true|none|none|
|»»» writers|[string]|true|none|none|
|»»» characters|[string]|true|none|none|
|»»» type|[string]|true|none|none|
|»»» region|string|true|none|none|
|»»» language|string|true|none|none|
|»»» release|string|true|none|none|
|»»» period|integer|true|none|none|
|»»» IMDb|string|true|none|none|
|»» score|object|true|none|none|
|»»» score|string|true|none|none|
|»»» total_cnt|integer|true|none|none|
|»»» five|string|true|none|none|
|»»» four|string|true|none|none|
|»»» three|string|true|none|none|
|»»» two|string|true|none|none|
|»»» one|string|true|none|none|
|»» plot|string|true|none|none|

## GET 获取电影的影评列表

GET /subjects/{id}/reviews

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|start|query|string|false|开始序列号，不填默认为0|
|limit|query|string|false|数量限制，不填为20|
|sort|query|string|false|排序规则｜填 hotest(最热门)|latest(最新)，不填为hotest|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取电影的讨论列表

GET /subjects/{id}/discussions

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|start|query|string|false|开始序列号，不填默认为0|
|limit|query|string|false|数量限制，不填为20|
|sort|query|string|false|排序规则｜填 hotest(最热门)|latest(最新)，不填为hotest|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 用户/获取信息

## GET 获取用户的看过短评列表

GET /users/{id}/after

+ 获取从 start 开始后 ≤ limit 个短评列表

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|start|query|string|false|开始序列号，不填默认为0|
|limit|query|string|false|数量限制，不填为20|
|sort|query|string|false|排序规则｜填 hotest(最热门)|latest(最新)，不填为latest|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取用户的影评

GET /users/{id}/reviews

+ 返回从 start 开始的 ≤ limit 个影评快照列表

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|start|query|string|false|开始序列号，不填默认为0|
|limit|query|string|false|数量限制，不填为20|
|sort|query|string|false|排序规则｜填 hotest(最热门)|latest(最新)，不填为latest|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取用户的主页信息

GET /users/{id}

+ 请求中 scope 有什么，返回的 json 的 data 里面就有什么字段
+ 不建议一次性把 scope 填满
+ 没有 scope 视为空字符串，只返回基础信息
+ 返回的 scope 数据有数量限制，详情看数据模型

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|scope|query|string|false|请求的范围｜从reviews|before|after中选取，多个用 , 隔开|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|获取基础信息(scope为空)|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» status|integer|true|none|none|
|» info|string|true|none|none|
|» data|object|true|none|none|
|»» username|string|true|none|none|
|»» uid|integer|true|none|none|
|»» github_id|integer|true|none|none|
|»» gitee_id|integer|true|none|none|
|»» email|string|true|none|none|
|»» phone|string|true|none|none|
|»» avatar|string|true|none|none|

## GET 获取用户片单列表

GET /users/{id}/movie_list

+ 返回从 start 开始后的 ≤limit 个片单列表信息

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|start|query|string|false|开始序列号，不填默认为0|
|limit|query|string|false|数量限制，不填为20|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取用户的想看短评列表

GET /users/{id}/before

+ 获取从 start 开始后 ≤ limit 个短评列表

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|start|query|string|false|开始序列号，不填默认为0|
|limit|query|string|false|数量限制，不填为20|
|sort|query|string|false|排序规则｜填 hotest(最热门)|latest(最新)，不填为latest|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 用户/更新信息

## PUT 更新用户非重要信息

PUT /users/{id}

+ 更新用户非重要信息
+ 需要验证
+ 验证方法
 Authorization jwt 认证
+ scope：每填入一个就需要在form_data里加一个

> Body 请求参数

```yaml
type: object
properties:
  scope:
    type: string
    description: 更新范围｜从username|github_id|gitee_id|avatar|description 中选取，多个用 , 隔开
  username:
    type: string
  github_id:
    type: string
  wechat_id:
    type: string
  avatar:
    type: string
  description:
    type: string
required:
  - scope

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|
|body|body|object|false|none|
|» scope|body|string|true|更新范围｜从username|github_id|gitee_id|avatar|description 中选取，多个用 , 隔开|
|» username|body|string|false|none|
|» github_id|body|string|false|none|
|» wechat_id|body|string|false|none|
|» avatar|body|string|false|none|
|» description|body|string|false|none|

> 返回示例

> 成功

```json
{
  "status": 20005,
  "info": "success",
  "data": {
    "detail": "operation success"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» status|integer|true|none|none|
|» info|string|true|none|none|
|» data|object|true|none|none|
|»» detail|string|true|none|none|

## PATCH 更新用户重要信息

PATCH /users/{id}

+ 修改重要信息
+ verify 验证码必填，需要先调用 给邮箱|手机号发送验证码 接口来发送验证码，然后才能调用本接口
+ 不能使用邮箱认证后修改成其他的邮箱，例如用 xxx@a.com 认证用来修改成 xxx@b.com，最终会被认证的邮箱覆盖
+ 短信认证同上
+ 验证码验证并非证实身份，而是绑定一个账户信息

> Body 请求参数

```yaml
type: object
properties:
  scope:
    type: string
    description: 范围｜从 password|email|phone 中选择，多个用 , 隔开
  verify_account:
    type: string
    description: 必填｜验证码账号
  verify:
    type: string
    description: 必填｜验证码
  verify_type:
    type: string
    description: 必填｜验证码种类
  password:
    type: string
  email:
    type: string
  phone:
    type: string
required:
  - scope
  - verify_account
  - verify
  - verify_type

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|
|body|body|object|false|none|
|» scope|body|string|true|范围｜从 password|email|phone 中选择，多个用 , 隔开|
|» verify_account|body|string|true|必填｜验证码账号|
|» verify|body|string|true|必填｜验证码|
|» verify_type|body|string|true|必填｜验证码种类|
|» password|body|string|false|none|
|» email|body|string|false|none|
|» phone|body|string|false|none|

> 返回示例

> 成功

```json
{
  "status": 20005,
  "info": "success",
  "data": {
    "detail": "operation success"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» status|integer|true|none|none|
|» info|string|true|none|none|
|» data|object|true|none|none|
|»» detail|string|true|none|none|

# 用户/账户

## POST 用户注册

POST /users/register

+ 邮箱和用户名登录使用密码
+ 手机号用短信验证登录

> Body 请求参数

```yaml
type: object
properties:
  account:
    type: string
    description: 账户｜手机号｜用户名｜电子邮箱
  token:
    type: string
    description: 令牌｜密码｜验证码
  type:
    type: string
    description: 方式｜password(密码式)｜email(邮箱式)｜sms(短信式)
    example: password
required:
  - account
  - token
  - type

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» account|body|string|true|账户｜手机号｜用户名｜电子邮箱|
|» token|body|string|true|令牌｜密码｜验证码|
|» type|body|string|true|方式｜password(密码式)｜email(邮箱式)｜sms(短信式)|

> 返回示例

> 成功

```json
{
  "status": 20000,
  "info": "success",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
    "uid": 43
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» status|integer|true|none|none|
|» info|string|true|none|none|
|» data|object|true|none|none|
|»» access_token|string|true|none|none|
|»» refresh_token|string|true|none|none|
|»» uid|integer|true|none|none|

## GET OAuth登录

GET /users/login

+ 第三方登录
+ 链接跳转
+ 预计支持 github 和 gitee

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|platform|query|string|true|平台｜github｜gitee|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 用户登录

POST /users/login

+ 当使用 refresh (刷新令牌) 时可以不用 account

> Body 请求参数

```yaml
type: object
properties:
  account:
    type: string
    description: 账户｜手机号｜用户名｜电子邮箱
  token:
    type: string
    description: 令牌｜密码｜验证码｜刷新令牌
  type:
    type: string
    description: 方式｜password(密码式)｜email(邮箱式)｜sms(短信式)｜refresh(刷新令牌式)
    example: password
required:
  - account
  - token
  - type

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» account|body|string|true|账户｜手机号｜用户名｜电子邮箱|
|» token|body|string|true|令牌｜密码｜验证码｜刷新令牌|
|» type|body|string|true|方式｜password(密码式)｜email(邮箱式)｜sms(短信式)｜refresh(刷新令牌式)|

> 返回示例

> 成功

```json
{
  "status": 20000,
  "info": "success",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
    "uid": 98
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» status|integer|true|none|none|
|» info|string|true|none|none|
|» data|object|true|none|none|
|»» access_token|string|true|none|none|
|»» refresh_token|string|true|none|none|
|»» uid|integer|true|none|none|

## GET 给确定用户发送验证码

GET /users/{id}/verify

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|type|query|string|true|验证码类型｜sms｜email|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 销毁账户

DELETE /users/{id}

+ 调用前使用 给确定用户发验证码 这个接口发送验证码
+ 需要 jwt 认证 + 验证码认证

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 给邮箱|手机号发送验证码

GET /users/verify

+ 两次验证码发送请求时间不能短于60s
+ 验证码会在服务器保存到下一次验证码发来的时候

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|type|query|string|true|验证码种类｜sms(短信验证)｜email(邮箱验证)|
|value|query|string|true|值｜手机号码｜电子邮箱|

> 返回示例

> 发送email成功

```json
{
  "status": 20002,
  "info": "success",
  "data": {
    "detail": "sending email success"
  }
}
```

> 验证码重发间隔太短

```json
{
  "status": 40000,
  "info": "invalid sending",
  "data": {
    "detail": "sending period too short"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|发送email成功|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|验证码重发间隔太短|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» status|integer|true|none|none|
|» info|string|true|none|none|
|» data|object|true|none|none|
|»» detail|string|true|none|none|

状态码 **400**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» status|integer|true|none|none|
|» info|string|true|none|none|
|» data|object|true|none|none|
|»» detail|string|true|none|none|

## POST 忘记密码/重置密码

POST /users/forget

+ 调用本接口前需要先调用给用户发送验证码接口
+ 用于忘记密码
+ 短信认证和邮箱认证
+ 不允许有 Authorization 字段
+ 当用户无法登录才能使用本接口

> Body 请求参数

```yaml
type: object
properties:
  uid:
    type: string
    description: 用户 uid
  verify:
    type: string
    description: 验证码
  verify_type:
    type: string
    description: 验证码类型
  new_pwd:
    type: string
    description: 重置的密码
required:
  - uid
  - verify
  - verify_type
  - new_pwd

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» uid|body|string|true|用户 uid|
|» verify|body|string|true|验证码|
|» verify_type|body|string|true|验证码类型|
|» new_pwd|body|string|true|重置的密码|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 用户/操作

## PATCH 关注

PATCH /users/{id}/following

+ 需要鉴权

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|id|query|string|true|关注的id|
|type|query|string|true|关注类型｜填 users|lists 表示关注用户或者片单|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 取关

DELETE /users/{id}/following

+ 需要鉴权

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|id|query|string|true|取关的id|
|type|query|string|true|取关类型｜填 users|lists 表示取关用户或者片单|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 片单

## GET 获取片单信息

GET /lists/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 更新片单信息

PUT /lists/{id}

+ 只能更新除影片列表以外信息
+ 需要鉴权

> Body 请求参数

```yaml
type: object
properties:
  scope:
    type: string
    description: 范围｜填 name|description，多个用 , 隔开
  name:
    type: string
  description:
    type: string
required:
  - scope

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|
|body|body|object|false|none|
|» scope|body|string|true|范围｜填 name|description，多个用 , 隔开|
|» name|body|string|false|none|
|» description|body|string|false|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除片单

DELETE /lists/{id}

+ 需要鉴权

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 创建片单

POST /lists

+ 需要 jwt 认证身份

> Body 请求参数

```yaml
type: object
properties:
  name:
    type: string
    description: 片单名称
  list:
    type: string
    description: 电影id列表，用 , 隔开多个
  description:
    type: string
    description: 片单描述

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string|true|JWT 鉴权|
|body|body|object|false|none|
|» name|body|string|false|片单名称|
|» list|body|string|false|电影id列表，用 , 隔开多个|
|» description|body|string|false|片单描述|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PATCH 向片单添加电影

PATCH /lists/{id}/movie

+ 需要鉴权

> Body 请求参数

```yaml
type: object
properties:
  mid:
    type: string
    description: 电影id列表，多个用 , 隔开
required:
  - mid

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|
|body|body|object|false|none|
|» mid|body|string|true|电影id列表，多个用 , 隔开|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 片单内删除电影

DELETE /lists/{id}/movie

+ 需要鉴权

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 影人

## GET 获取影人信息

GET /celebrities/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 短评|想看|看过

## POST 发布短评

POST /comments

+ 需要鉴权

> Body 请求参数

```yaml
type: object
properties:
  mid:
    type: string
    description: 电影id
  score:
    type: string
    description: 评分 1-5
  type:
    type: string
    description: 填 before|after
  content:
    type: string
    description: 内容
  tag:
    type: string
    description: 标签，用 , 隔开多个
required:
  - mid
  - score
  - type
  - content
  - tag

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string|true|JWT 鉴权|
|body|body|object|false|none|
|» mid|body|string|true|电影id|
|» score|body|string|true|评分 1-5|
|» type|body|string|true|填 before|after|
|» content|body|string|true|内容|
|» tag|body|string|true|标签，用 , 隔开多个|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除短评

DELETE /comments/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 更新短评信息

PUT /comments/{id}

> Body 请求参数

```yaml
type: object
properties:
  score:
    type: string
    description: 1-5
  tag:
    type: string
    description: 标签，用 , 隔开多个
  content:
    type: string
    description: 内容
required:
  - score
  - tag
  - content

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|
|body|body|object|false|none|
|» score|body|string|true|1-5|
|» tag|body|string|true|标签，用 , 隔开多个|
|» content|body|string|true|内容|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取短评信息

GET /comments/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 点赞|取消点赞

GET /comments/{id}/star

+ 点赞
+ 会写入一个 cookie 存储点赞关系
+ 需要鉴权

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|type|query|string|true|类型|
|value|query|string|true|true|fase  点赞|取消点赞|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 讨论

## POST 发布讨论

POST /discussions

> Body 请求参数

```yaml
type: object
properties:
  mid:
    type: string
    description: 电影id
  name:
    type: string
    description: 标题
  content:
    type: string
    description: 内容
required:
  - mid
  - name
  - content

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string|true|JWT 鉴权|
|body|body|object|false|none|
|» mid|body|string|true|电影id|
|» name|body|string|true|标题|
|» content|body|string|true|内容|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取讨论信息

GET /discussions/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 更新讨论

PUT /discussions/{id}

> Body 请求参数

```yaml
type: object
properties:
  name:
    type: string
    description: 标题
  content:
    type: string
    description: 内容
required:
  - name
  - content

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|
|body|body|object|false|none|
|» name|body|string|true|标题|
|» content|body|string|true|内容|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除讨论

DELETE /discussions/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 点赞|取消点赞

GET /discussions/{id}/star

+ 同短评的点赞

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|value|query|string|true|true|false，点赞|取消点赞|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 影评|长评

## GET 获取影评信息

GET /reviews/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 更新影评内容

PUT /reviews/{id}

> Body 请求参数

```yaml
type: object
properties:
  score:
    type: string
    description: 评分 1-5
  name:
    type: string
    description: 标题
  content:
    type: string
    description: 内容
required:
  - score
  - name
  - content

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|
|body|body|object|false|none|
|» score|body|string|true|评分 1-5|
|» name|body|string|true|标题|
|» content|body|string|true|内容|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除影评

DELETE /reviews/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 点赞|取消点赞

GET /reviews/{id}/star

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|review_star_list|cookie|string|true|已经点过赞的|
|id|path|string|true|none|
|value|query|string|true|true|false|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 发布影评

POST /reviews

> Body 请求参数

```yaml
type: object
properties:
  mid:
    type: string
    description: 电影id
  score:
    type: string
    description: 评分 1-5
  name:
    type: string
    description: 标题
  content:
    type: string
    description: 内容
required:
  - mid
  - score
  - name
  - content

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string|true|JWT 鉴权|
|body|body|object|false|none|
|» mid|body|string|true|电影id|
|» score|body|string|true|评分 1-5|
|» name|body|string|true|标题|
|» content|body|string|true|内容|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 点踩|取消点踩

GET /reviews/{id}/bad

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|value|query|string|true|true|false|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 其他API

## POST 上传头像

POST /avatar

+ 需要鉴权
+ 将头像上传到储存桶里，返回一个链接，并且会将头像设置给用户

> Body 请求参数

```yaml
type: object
properties:
  img:
    type: string
    description: 图像
    format: binary
required:
  - img

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string|true|JWT 鉴权|
|body|body|object|false|none|
|» img|body|string(binary)|true|图像|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 外链跳转

GET /wild

+ 为了预防 CSRF 攻击，所有外链跳转都使用本接口
+ 注意：请求头里不能带有 Authorization 字段，否则拒绝跳转

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|link|query|string|true|跳转的外链链接|

> 返回示例

> 请求有误

```json
{
  "status": 40004,
  "info": "invalid request",
  "data": {
    "detail": "can not go wild"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|请求有误|Inline|

### 返回数据结构

状态码 **400**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» status|integer|true|none|不允许有 Authorization|
|» info|string|true|none|none|
|» data|object|true|none|none|
|»» detail|string|true|none|none|

## GET 获取OpenAPI文档

GET /swagger

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 搜索

GET /search

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|key|query|string|true|搜索关键词|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 我的主页

GET /mine

+ 等同访问 /users/{id}
+ 需要 Authorization 进行身份认证识别

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 回复|回应

## POST 发表回复

POST /replies

> Body 请求参数

```yaml
type: object
properties:
  pid:
    type: string
    description: 父id
  type:
    type: string
    description: 类型｜填 review|discussion|comment|reply
  content:
    type: string
    description: 内容
required:
  - pid
  - type
  - content

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string|true|JWT 鉴权|
|body|body|object|false|none|
|» pid|body|string|true|父id|
|» type|body|string|true|类型｜填 review|discussion|comment|reply|
|» content|body|string|true|内容|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除回复

DELETE /replies/{id}

+ 会删掉这个评论和所有子评论

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|Authorization|header|string|true|JWT 鉴权|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 查询某个帖子回复

GET /replies/{id}

+ 只返回一级回复

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|帖子id|
|type|query|string|true|帖子类型｜填 review|discussion|comment|reply|
|limit|query|string|true|限制|
|start|query|string|true|offset|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 查询一个回复的所有子回复

GET /replies/{id}/all

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

<h2 id="tocS_电影主页">电影主页</h2>
<!-- backwards compatibility -->
<a id="schema电影主页"></a>
<a id="schema_电影主页"></a>
<a id="tocS电影主页"></a>
<a id="tocs电影主页"></a>

```json
{
  "type": "object",
  "properties": {
    "id": {
      "type": "integer",
      "title": "电影id"
    },
    "name": {
      "type": "string",
      "title": "电影名称"
    },
    "stars": {
      "type": "string",
      "title": "平均星星数"
    },
    "score": {
      "type": "object",
      "properties": {
        "total": {
          "type": "number",
          "title": "总评分"
        },
        "total_cnt": {
          "type": "integer",
          "title": "总评分人数"
        },
        "five": {
          "type": "number",
          "title": "五星比例"
        },
        "four": {
          "type": "number",
          "title": "四星比例"
        },
        "three": {
          "type": "number",
          "title": "三星比例"
        },
        "two": {
          "type": "number",
          "title": "二星比例"
        },
        "one": {
          "type": "number",
          "title": "一星比例"
        }
      },
      "required": [
        "total",
        "total_cnt",
        "five",
        "four",
        "three",
        "two",
        "one"
      ],
      "title": "评分信息"
    },
    "detail": {
      "type": "object",
      "properties": {
        "nicknames": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "电影的又名"
        },
        "director": {
          "type": "string",
          "title": "导演",
          "description": "多个导演用,隔开"
        },
        "writers": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "编剧"
        },
        "characters": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "主演"
        },
        "type": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "类型"
        },
        "website": {
          "type": "string",
          "title": "官网"
        },
        "region": {
          "type": "string",
          "title": "制片地区/国家"
        },
        "language": {
          "type": "string",
          "title": "语言"
        },
        "release": {
          "type": "string",
          "title": "上映时间",
          "description": "中国大陆时间"
        },
        "period": {
          "type": "integer",
          "title": "片长",
          "description": "单位：分钟"
        },
        "IMDb": {
          "type": "string",
          "title": "IMDb"
        }
      },
      "title": "电影的基础信息",
      "required": [
        "nicknames",
        "director",
        "writers",
        "characters",
        "type",
        "website",
        "region",
        "language",
        "release",
        "period",
        "IMDb"
      ]
    },
    "plot": {
      "type": "string",
      "title": "剧情简介"
    },
    "celebrities": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "id": {
            "type": "integer",
            "description": "影人的id"
          },
          "name": {
            "type": "string",
            "description": "影人的名称"
          },
          "job": {
            "type": "string",
            "description": "影人在此剧的职位，多个职位用,隔开"
          },
          "avatar": {
            "type": "string",
            "description": "影人的头像"
          }
        },
        "description": "影人的快照",
        "required": [
          "id",
          "name",
          "job",
          "avatar"
        ]
      },
      "title": "影人",
      "description": "限制最多6个"
    },
    "comments": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "tag": {
            "type": "array",
            "items": {
              "type": "string",
              "description": "标签"
            },
            "title": "标签列表"
          },
          "id": {
            "type": "integer",
            "title": "短评id"
          },
          "mid": {
            "type": "integer",
            "title": "电影id"
          },
          "content": {
            "type": "string",
            "title": "内容",
            "description": "不超过165个字"
          },
          "score": {
            "type": "integer",
            "title": "评分",
            "description": "不能超过5"
          },
          "username": {
            "type": "string",
            "title": "用户名称"
          },
          "uid": {
            "type": "string",
            "title": "用户id"
          },
          "type": {
            "type": "string",
            "title": "分类｜before｜after"
          },
          "date": {
            "type": "string",
            "title": "日期"
          },
          "stars": {
            "type": "integer",
            "title": "点赞数"
          }
        },
        "required": [
          "tag",
          "content",
          "score",
          "username",
          "uid",
          "type",
          "date",
          "stars",
          "id",
          "mid"
        ],
        "x-apifox-folder": ""
      },
      "title": "热门短评",
      "description": "限制最多5个(前五的点赞数)"
    },
    "reviews": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {}
      },
      "title": "热门影评",
      "description": "限制最多10个(前10)"
    },
    "discussions": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string",
            "title": "短评名称"
          },
          "id": {
            "type": "integer",
            "title": "短评id"
          },
          "username": {
            "type": "string",
            "title": "短评用户名"
          },
          "uid": {
            "type": "integer",
            "title": "短评用户id"
          },
          "reply_cnt": {
            "type": "integer",
            "title": "短评回复数"
          },
          "date": {
            "type": "string",
            "title": "短评日期"
          }
        },
        "required": [
          "name",
          "id",
          "username",
          "uid",
          "reply_cnt",
          "date"
        ],
        "x-apifox-folder": ""
      },
      "title": "最新讨论",
      "description": "限制最多5个"
    }
  },
  "required": [
    "id",
    "name",
    "detail",
    "score",
    "plot",
    "celebrities",
    "comments",
    "reviews",
    "discussions",
    "stars"
  ],
  "x-apifox-folder": ""
}

```

### 属性

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|id|integer|true|none|none|
|name|string|true|none|none|
|stars|string|true|none|none|
|score|object|true|none|none|
|» total|number|true|none|none|
|» total_cnt|integer|true|none|none|
|» five|number|true|none|none|
|» four|number|true|none|none|
|» three|number|true|none|none|
|» two|number|true|none|none|
|» one|number|true|none|none|
|detail|object|true|none|none|
|» nicknames|[string]|true|none|none|
|» director|string|true|none|多个导演用,隔开|
|» writers|[string]|true|none|none|
|» characters|[string]|true|none|none|
|» type|[string]|true|none|none|
|» website|string|true|none|none|
|» region|string|true|none|none|
|» language|string|true|none|none|
|» release|string|true|none|中国大陆时间|
|» period|integer|true|none|单位：分钟|
|» IMDb|string|true|none|none|
|plot|string|true|none|none|
|celebrities|[object]|true|none|限制最多6个|
|» id|integer|true|none|影人的id|
|» name|string|true|none|影人的名称|
|» job|string|true|none|影人在此剧的职位，多个职位用,隔开|
|» avatar|string|true|none|影人的头像|
|comments|[[%E7%9F%AD%E8%AF%84](#schema%e7%9f%ad%e8%af%84)]|true|none|限制最多5个(前五的点赞数)|
|reviews|[object]|true|none|限制最多10个(前10)|
|discussions|[[%E8%AE%A8%E8%AE%BA%E5%BF%AB%E7%85%A7](#schema%e8%ae%a8%e8%ae%ba%e5%bf%ab%e7%85%a7)]|true|none|限制最多5个|

<h2 id="tocS_用户主页">用户主页</h2>
<!-- backwards compatibility -->
<a id="schema用户主页"></a>
<a id="schema_用户主页"></a>
<a id="tocS用户主页"></a>
<a id="tocs用户主页"></a>

```json
{
  "type": "object",
  "properties": {
    "username": {
      "type": "string",
      "title": "用户名称"
    },
    "password": {
      "type": "string",
      "title": "用户密码",
      "description": "不会作为返回出现"
    },
    "uid": {
      "type": "integer",
      "title": "用户id"
    },
    "github_id": {
      "type": "integer",
      "title": "用户的github账户"
    },
    "gitee_id": {
      "type": "integer",
      "title": "用户的gitee账户"
    },
    "email": {
      "type": "string",
      "title": "用户的邮箱"
    },
    "phone": {
      "type": "string",
      "title": "用户的手机号码"
    },
    "avatar": {
      "type": "string",
      "title": "用户头像url"
    },
    "reviews": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {}
      },
      "title": "影评列表",
      "description": "限制4个"
    },
    "before": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "tag": {
            "type": "array",
            "items": {
              "type": "string",
              "description": "标签"
            },
            "title": "标签列表"
          },
          "id": {
            "type": "integer",
            "title": "短评id"
          },
          "mid": {
            "type": "integer",
            "title": "电影id"
          },
          "content": {
            "type": "string",
            "title": "内容",
            "description": "不超过165个字"
          },
          "score": {
            "type": "integer",
            "title": "评分",
            "description": "不能超过5"
          },
          "username": {
            "type": "string",
            "title": "用户名称"
          },
          "uid": {
            "type": "string",
            "title": "用户id"
          },
          "type": {
            "type": "string",
            "title": "分类｜before｜after"
          },
          "date": {
            "type": "string",
            "title": "日期"
          },
          "stars": {
            "type": "integer",
            "title": "点赞数"
          }
        },
        "required": [
          "tag",
          "content",
          "score",
          "username",
          "uid",
          "type",
          "date",
          "stars",
          "id",
          "mid"
        ],
        "x-apifox-folder": ""
      },
      "title": "用户的想看",
      "description": "限制10个"
    },
    "after": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "tag": {
            "type": "array",
            "items": {
              "type": "string",
              "description": "标签"
            },
            "title": "标签列表"
          },
          "id": {
            "type": "integer",
            "title": "短评id"
          },
          "mid": {
            "type": "integer",
            "title": "电影id"
          },
          "content": {
            "type": "string",
            "title": "内容",
            "description": "不超过165个字"
          },
          "score": {
            "type": "integer",
            "title": "评分",
            "description": "不能超过5"
          },
          "username": {
            "type": "string",
            "title": "用户名称"
          },
          "uid": {
            "type": "string",
            "title": "用户id"
          },
          "type": {
            "type": "string",
            "title": "分类｜before｜after"
          },
          "date": {
            "type": "string",
            "title": "日期"
          },
          "stars": {
            "type": "integer",
            "title": "点赞数"
          }
        },
        "required": [
          "tag",
          "content",
          "score",
          "username",
          "uid",
          "type",
          "date",
          "stars",
          "id",
          "mid"
        ],
        "x-apifox-folder": ""
      },
      "title": "用户的看过",
      "description": "限制10个"
    }
  },
  "required": [
    "username",
    "uid",
    "github_id",
    "gitee_id",
    "email",
    "phone",
    "avatar",
    "reviews",
    "before",
    "after",
    "password"
  ],
  "x-apifox-folder": ""
}

```

### 属性

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|username|string|true|none|none|
|password|string|true|none|不会作为返回出现|
|uid|integer|true|none|none|
|github_id|integer|true|none|none|
|gitee_id|integer|true|none|none|
|email|string|true|none|none|
|phone|string|true|none|none|
|avatar|string|true|none|none|
|reviews|[object]|true|none|限制4个|
|before|[[%E7%9F%AD%E8%AF%84](#schema%e7%9f%ad%e8%af%84)]|true|none|限制10个|
|after|[[%E7%9F%AD%E8%AF%84](#schema%e7%9f%ad%e8%af%84)]|true|none|限制10个|

<h2 id="tocS_片单">片单</h2>
<!-- backwards compatibility -->
<a id="schema片单"></a>
<a id="schema_片单"></a>
<a id="tocS片单"></a>
<a id="tocs片单"></a>

```json
{
  "type": "object",
  "properties": {
    "id": {
      "type": "integer"
    },
    "uid": {
      "type": "integer"
    },
    "date": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "followers": {
      "type": "array",
      "items": {
        "type": "integer"
      }
    },
    "list": {
      "type": "array",
      "items": {
        "type": "integer"
      }
    },
    "description": {
      "type": "string"
    }
  },
  "required": [
    "id",
    "uid",
    "date",
    "name",
    "followers",
    "list",
    "description"
  ],
  "x-apifox-folder": ""
}

```

### 属性

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|id|integer|true|none|none|
|uid|integer|true|none|none|
|date|string|true|none|none|
|name|string|true|none|none|
|followers|[integer]|true|none|none|
|list|[integer]|true|none|none|
|description|string|true|none|none|

<h2 id="tocS_讨论快照">讨论快照</h2>
<!-- backwards compatibility -->
<a id="schema讨论快照"></a>
<a id="schema_讨论快照"></a>
<a id="tocS讨论快照"></a>
<a id="tocs讨论快照"></a>

```json
{
  "type": "object",
  "properties": {
    "name": {
      "type": "string",
      "title": "短评名称"
    },
    "id": {
      "type": "integer",
      "title": "短评id"
    },
    "username": {
      "type": "string",
      "title": "短评用户名"
    },
    "uid": {
      "type": "integer",
      "title": "短评用户id"
    },
    "reply_cnt": {
      "type": "integer",
      "title": "短评回复数"
    },
    "date": {
      "type": "string",
      "title": "短评日期"
    }
  },
  "required": [
    "name",
    "id",
    "username",
    "uid",
    "reply_cnt",
    "date"
  ],
  "x-apifox-folder": ""
}

```

### 属性

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|name|string|true|none|none|
|id|integer|true|none|none|
|username|string|true|none|none|
|uid|integer|true|none|none|
|reply_cnt|integer|true|none|none|
|date|string|true|none|none|

<h2 id="tocS_短评">短评</h2>
<!-- backwards compatibility -->
<a id="schema短评"></a>
<a id="schema_短评"></a>
<a id="tocS短评"></a>
<a id="tocs短评"></a>

```json
{
  "type": "object",
  "properties": {
    "tag": {
      "type": "array",
      "items": {
        "type": "string",
        "description": "标签"
      },
      "title": "标签列表"
    },
    "id": {
      "type": "integer",
      "title": "短评id"
    },
    "mid": {
      "type": "integer",
      "title": "电影id"
    },
    "content": {
      "type": "string",
      "title": "内容",
      "description": "不超过165个字"
    },
    "score": {
      "type": "integer",
      "title": "评分",
      "description": "不能超过5"
    },
    "username": {
      "type": "string",
      "title": "用户名称"
    },
    "uid": {
      "type": "string",
      "title": "用户id"
    },
    "type": {
      "type": "string",
      "title": "分类｜before｜after"
    },
    "date": {
      "type": "string",
      "title": "日期"
    },
    "stars": {
      "type": "integer",
      "title": "点赞数"
    }
  },
  "required": [
    "tag",
    "content",
    "score",
    "username",
    "uid",
    "type",
    "date",
    "stars",
    "id",
    "mid"
  ],
  "x-apifox-folder": ""
}

```

### 属性

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|tag|[string]|true|none|none|
|id|integer|true|none|none|
|mid|integer|true|none|none|
|content|string|true|none|不超过165个字|
|score|integer|true|none|不能超过5|
|username|string|true|none|none|
|uid|string|true|none|none|
|type|string|true|none|none|
|date|string|true|none|none|
|stars|integer|true|none|none|

