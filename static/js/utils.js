import { ACCESS_TOKEN, BASE_URL } from './consts.js'

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
    const authorization = localStorage.getItem(ACCESS_TOKEN)
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
        for (let k in params) {
            if (!url.endsWith("?")) {
                url += "&"
            }
            url += `${k}=${params[k]}`
        }
    }
    const authorization = localStorage.getItem(ACCESS_TOKEN)
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

// PUT请求
async function put(path, body, params) {
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
    const authorization = localStorage.getItem(ACCESS_TOKEN)
    if (authorization != null) {
        return await (await fetch(url, {
            method: "PUT",
            body,
            headers: {
                'Authorization': authorization
            },
        })).json()
    }
    return await (await fetch(url, {
        method: "PUT",
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

/**
 * 获取token中含有的用户ID
 * 
 * @param {string} access_token 
 * @returns 
 */
function getUserIdFromToken(access_token) {
    return getTokenInfoObj(access_token).uid
}

function getUserId() {
    return getUserIdFromToken(localStorage.getItem(ACCESS_TOKEN))
}

// 获取绝对地址
function getAbsolutePath(path) {
    const curWwwPath = window.document.location.href
    const pathName = window.document.location.pathname;
    const pos = curWwwPath.indexOf(pathName)
    const localhostPaht = curWwwPath.substring(0, pos)
    return localhostPaht + path
}

export { get, post, getTokenInfoObj, put, getUserIdFromToken, getUserId, getAbsolutePath }