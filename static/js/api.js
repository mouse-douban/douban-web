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
 * @param {int} id 
 * @param {string} username 
 * @param {string} avatar 
 * @param {string} scope 更新范围｜从username|github_id|gitee_id|avatar 中选取，多个用 , 隔开
 */
async function putUserInfo(id, username, avatar, scope = "username, avatar") {
    const data = new FormData()
    data.append("username", username)
    data.append("avatar", avatar)
    return await put(`/users/${id}`, data, { scope })
}

export { getUserInfo, login, getVerifyCode, getMineInfo, putUserInfo }