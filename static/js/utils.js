import { BASE_URL } from './consts.js'

// GET请求
async function get(path, params = null) {
    let url = BASE_URL + path
    if (params != null) {
        url += "?"
        for (let k in params) {
            if (!url.endsWith("?")) {
                url += "&"
            }
            url += `${k}=${params[k]}`
        }
    }
    const authorization = localStorage.getItem("access_token")
    if (authorization != null) {
        return await (await fetch(url, {
            method: "GET",
            headers: {
                'Authorization': authorization
            },
        })).json()
    }
    return await (await fetch(url, {
        method: "GET",
    })).json()
}

// POST请求
async function post(path, body, params = null) {
    let url = BASE_URL + path
    if (params != null) {
        url += "?"
        for (k in params) {
            url += `${k}=${params[k]}&`
        }
        url -= "&"
    }
    const authorization = localStorage.getItem("authorization")
    if (authorization != null) {
        return await (await fetch(url, {
            method: "POST",
            body,
            headers: {
                'Authorization': authorization
            },
        })).json()
    }
    return await (await fetch(url, {
        method: "POST",
        body,
    })).json()
}

/**
 * 获取token中含有的基础信息 (Token有效期 签发时间 用户id)
 * 
 * 为什么atob才是解码啊？？？？
 * 
 * @param access_token 
 * @returns 
 */
function getTokenInfoObj(access_token) {
    return JSON.parse(window.atob(access_token.split(".")[1]))
}

export { get, post, getTokenInfoObj }