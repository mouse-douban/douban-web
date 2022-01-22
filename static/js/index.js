import { getUserInfo, login } from './api.js'
import { getTokenInfoObj } from './utils.js'
import { ACCESS_TOKEN, REFRESH_TOKEN, USER_INFO } from './consts.js'

const loginProfileElement = document.querySelector("#login-nav-item")
let access_token = localStorage.getItem(ACCESS_TOKEN)
const refresh_token = localStorage.getItem(REFRESH_TOKEN)

// 自动登录部分
if (access_token != null) {
    const tokenLifeCycleObj = getTokenInfoObj(access_token)
    const isTokenAvaliable = new Date().getTime() - tokenLifeCycleObj.iat < tokenLifeCycleObj.exp
    // token有效 获取用户信息
    if (isTokenAvaliable) {
        doGetUserInfo()
    } else {
        // token无效 用refresh_token重新获取access_token
        // 获取成功则获取用户信息，失败则设置手动登录按钮
        tryAutoLogin()
    }
} else if (refresh_token != null) {
    tryAutoLogin()
} else {
    // 从未登录过 设置手动登录按钮
    setUpLoginButton()
}

async function doGetUserInfo() {
    const info = await getUserInfo(tokenLifeCycleObj.uid)
    localStorage.setItem(USER_INFO, JSON.stringify(info))
    loginProfileElement.textContent = `${info.data.username}的账号`
}

function setUpLoginButton() {
    loginProfileElement.addEventListener('click', () => {
        window.location.href = "./html/login.html"
    })
}

async function tryAutoLogin() {
    const res = await login("", refresh_token, "refresh")
    switch (res.status) {
        // 成功
        case 20000: {
            access_token = res.data.access_token
            doGetUserInfo()
            break
        }
        // 失败
        default: {
            setUpLoginButton()
        }
    }
}
