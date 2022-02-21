import { get, post, put } from './utils.js'
// ApiService

/**
 * 获取用户信息
 * 
 * @param {number} id 用户id
 * @param {string} scope 需求的详细信息 从reviews|movie_list|before|after中选取，多个用 , 隔开
 * @returns 
 */
async function getUserInfo(id, scope = "") {
    return await get(`/users/${id}`, { scope })
}

/**
 * 登录
 * 
 * @param {string} account 账户｜手机号｜用户名｜电子邮箱
 * @param {string} token 令牌｜密码｜验证码｜刷新令牌
 * @param {string} type 方式｜password(密码式)｜email(邮箱式)｜sms(短信式)｜refresh(刷新令牌式)
 */
async function login(account, token, type) {
    const formData = new FormData()
    formData.append("account", account)
    formData.append("token", token)
    formData.append("type", type)
    return await post("/users/login", formData)
}

/**
 * 获取验证码
 * 
 * @param {string} type 验证码种类 sms(短信) / email(邮件)
 * @param {string} value 值 电话号码 / 邮箱
 */
async function getVerifyCode(type, value) {
    return await get("/users/verify", {
        type,
        value
    })
}

/**
 * 获取自身信息，需要jwt鉴权
 * 
 * @param {string} scope 需求的详细信息 从reviews|movie_list|before|after中选取，多个用 , 隔开
 * @returns 
 */
async function getMineInfo(scope = "") {
    return await get("/mine", { scope })
}

/**
 * 修改不重要信息，需要jwt鉴权
 * 
 * @param {number} id 
 * @param {string} username 
 * @param {string} avatar 
 * @param {string} scope 更新范围｜从username|github_id|gitee_id|avatar 中选取，多个用 , 隔开
 */
async function putUserInfo(id, username, description, scope = "username,description") {
    const data = new FormData()
    data.append("scope", scope)
    data.append("username", username)
    data.append("description", description)
    return await put(`/users/${id}`, data)
}

/**
 * 获取用户想看列表
 * 
 * @param {number} id 
 * @param {number} start 开始序列号，不填默认为0
 * @param {number} limit 数量限制，不填为20
 * @param {string} sort 排序规则｜填 hotest(最热门)|latest(最新)，不填为latest
 * @returns 
 */
async function getWishToWatchList(id, start = 0, limit = 20, sort = "latest") {
    return await get(`/users/${id}/before`, { start, limit, sort })
}

/**
 * 获取用户看过列表
 * 
 * @param {number} id 
 * @param {number} start 开始序列号，不填默认为0
 * @param {number} limit 数量限制，不填为20
 * @param {string} sort 排序规则｜填 hotest(最热门)|latest(最新)，不填为latest
 * @returns 
 */
async function getWatchedList(id, start = 0, limit = 20, sort = "latest") {
    return await get(`/users/${id}/after`, { start, limit, sort })
}

/**
 * 获取电影信息
 * 
 * @param {number} id 电影id
 * @param {string} scope 请求的范围｜从plot|celebrities|comments|reviews|discussions 中选取，多个用 , 隔开
 * @returns 
 */
async function getMovieInfo(id, scope = "") {
    return await get(`/subjects/${id}`, { scope })
}

/**
 * 获取用户评论
 * 
 * @param {number} id 
 * @param {number} start 开始序列号，不填默认为0
 * @param {number} limit 数量限制，不填为20
 * @param {string} sort 排序规则｜填 hotest(最热门)|latest(最新)，不填为latest
 * @returns 
 */
async function getUserReviews(id, start = 0, limit = 20, sort = "latest") {
    return await get(`/users/${id}/reviews`, { start, limit, sort })
}

/**
 * 获取用户片单
 * 
 * @param {number} id 
 * @param {number} start 开始序列号，不填默认为0
 * @param {number} limit 数量限制，不填为20
 * @param {string} sort 排序规则｜填 hotest(最热门)|latest(最新)，不填为latest
 * @returns 
 */
async function getUserMovieList(id, start = 0, limit = 20, sort = "latest") {
    return await get(`/users/${id}/movie_list`, { start, limit, sort })
}

/**
 * 获取电影列表
 * 
 * @param {*} start 开始的索引，不填为0
 * @param {*} limit 限制数量，不填为20
 * @param {*} sort 排序规则｜填 hotest|latest，不填为 latest
 * @param {*} tag 标签 喜剧,生活,爱情,动作,科幻,悬疑,惊悚,动画,奇幻
 * @returns 
 */
async function getSubjects(start = 0, limit = 20, sort = "latest", tag = "") {
    return await get("/subjects", { start, limit, sort, tag })
}

/**
 * 获取电影主页信息
 * 
 * @param {*} id mid
 * @param {*} scope 范围, 填comments|reviews|discussions
 */
async function getSubjectById(id, scope = "") {
    return await get(`/subjects/${id}`, { scope })
}

/**
 * 获取影人信息
 * 
 * @param {*} id 
 * @returns 
 */
async function getCelebrityInfo(id) {
    return await get(`/celebrities/${id}`)
}

/**
 * 发布短评
 * 
 * @param {*} mid 
 * @param {*} score 
 * @param {*} type 
 * @param {*} content 
 * @param {*} tag 
 */
async function addComment(mid, score, type, content, tag) {
    const body = new FormData()
    body.append("mid", mid)
    body.append("score", score)
    body.append("type", type)
    body.append("content", content)
    body.append("tag", tag)
    return await post("/comments", body)
}

/**
 * 发布影评
 * 
 * @param {*} mid 
 * @param {*} score 
 * @param {*} name 
 * @param {*} content 
 */
async function addReviews(mid, score, name, content) {
    const body = new FormData()
    body.append("mid", mid)
    body.append("score", score)
    body.append("name", name)
    body.append("content", content)
    return await post("/reviews", body)
}

/**
 * 搜索
 * 
 * @param {*} key 
 */
async function search(key) {
    return await get("/search", { key })
}

/**
 * 忘记密码
 * 
 * @param {*} account 用户账号
 * @param {*} verify 验证码
 * @param {*} type 验证码类型
 * @param {*} new_pwd 新密码
 * @returns 
 */
async function forget(account, verify, type, new_pwd) {
    const body = new FormData()
    body.append("account", account)
    body.append("verify", verify)
    body.append("type", type)
    body.append("new_pwd", new_pwd)
    return await post("/users/forget", body)
}

/**
 * 获取讨论信息
 * 
 * @param {*} id 
 * @returns 
 */
async function getDiscussion(id) {
    return await get(`/discussions/${id}`)
}

/**
 * 点赞
 * 
 * @param {*} id 
 * @returns 
 */
async function starDiscussion(id, value = true) {
    return await get(`/discussions/${id}/star`, { value })
}

/**
 * 回复
 * 
 * @param {*} pid 父级id
 * @param {*} type review|discussion|comment|reply
 * @param {*} content 内容
 * @returns 
 */
async function reply(pid, type, content) {
    const body = new FormData()
    body.append("pid", pid)
    body.append("type", type)
    body.append("content", content)
    return post(`/replies`, body)
}

async function updateAvatar(avatar) {
    const body = new FormData()
    body.append("img", avatar)
    return await post("/avatar", body)
}

export {
    getUserInfo, login, getVerifyCode, getMineInfo, putUserInfo, getWatchedList, getWishToWatchList, getMovieInfo, getUserReviews,
    getUserMovieList, getSubjects, getSubjectById, getCelebrityInfo, addComment, addReviews, search, forget, getDiscussion,
    starDiscussion, reply, updateAvatar
}