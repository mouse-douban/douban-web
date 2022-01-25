import { get, post } from './utils.js'
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

export { getUserInfo, login }