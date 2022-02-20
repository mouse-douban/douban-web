import { BASE_URL } from "./consts.js"
import { forget, getVerifyCode } from "./api.js"

// switch state
const tabs = document.querySelectorAll('.account-tab')
const fragContainer = document.querySelector("#fragment-container")
// 选择短信登录/注册时显示的内容
const registerInnerHtml = `
                <p id="remind-content">请仔细阅读 <span style="color: #41ac52;cursor: pointer;" onclick="window.location.href = 'https://accounts.douban.com/passport/agreement'">豆瓣使用协议 豆瓣个人信息保护政策</span></p>
                <div class="input-box">
                    <div id="plus-86">+86</div>
                    <input id="input-phone" style="width: 80%;"
                        size="22" maxlength="60" placeholder="手机号">
                </div>
                <div class="input-box">
                    <input id="input-verification-code" style="width: 73%;"
                        size="22" maxlength="60" placeholder="输入验证码">
                    <div id="get-verification-code">获取验证码</div>
                </div>
                <div id="submit-btn">
                    登录豆瓣
                </div>
                <div style="width: 100%;height: fit-content;margin-top: 10px;display: flex;justify-content: end;">
                    <div id="cannot-get-verification-code" style="color:#41ac52;cursor: pointer;">收不到验证码</div>
                </div>
        `
// 选择密码登录时显示的内容
const loginInnerHtml = `
                <div style="margin-top: 20px;"></div>
                <div class="input-box">
                    <input id="input-id"
                        size="22" maxlength="60" placeholder="手机号/邮箱">
                </div>
                <div class="input-box">
                    <input id="input-password" type="password" style="width: 75%;"
                        size="22" maxlength="60" placeholder="输入密码">
                    <div id="get-verification-code" style="color: #9b9b9b; width: 25%;">找回密码</div>
                </div>
                <div id="submit-btn">
                    登录豆瓣
                </div>
        `
// 忘记密码的dialog
const dialogForgetPass = document.querySelector(".dialog-box")
const dialogAccount = document.querySelector("#account-input")
const dialogPassword = document.querySelector("#new-password-input")
const dialogVerifyCode = document.querySelector("#verify-code-input")
const smsRadio = document.querySelector("#sms")
// current state
let isRegisterView = true

// setup dialog
dialogForgetPass.querySelector(".dialog-closer").addEventListener("click", () => {
    dialogForgetPass.style.display = "none"
})

// submit
dialogForgetPass.querySelector("#submit-button").addEventListener("click", async () => {
    const type = smsRadio.checked ? "phone" : "email"
    const account = dialogAccount.value
    const vc = dialogVerifyCode.value
    const password = dialogPassword.value
    if (account.length === 0 || vc.length === 0 || password.length === 0) {
        alert("请填写完整信息")
        return
    }
    const data = await forget(account, vc, type, password)
    switch (data.status) {
        case 20000: {
            alert("操作成功！")
            dialogForgetPass.style.display = "none"
            break
        }
        default: {
            alert(data.data.detail)
        }
    }
})

// send verification code
dialogForgetPass.querySelector("#verify-code-btn").addEventListener("click", async () => {
    const type = smsRadio.checked ? "sms" : "email"
    const account = type == "sms" ? "%2B86" + dialogAccount.value : dialogAccount.value
    if (account.length === 0) {
        alert("请填写账号")
        return
    }
    const data = await getVerifyCode(type, account)
    switch (data.status) {
        case 20000: {
            alert("验证码已发送")
            break
        }
        default: {
            alert(data.data.detail)
        }
    }
})

tabs.forEach((value, key) => {
    value.addEventListener('click', _e => {
        if (!value.classList.contains("on")) {
            value.classList.add("on")
            const k = key == 0 ? 1 : 0
            if (value.textContent == "短信登录/注册") {
                fragContainer.innerHTML = registerInnerHtml
                isRegisterView = true
                setSubmitBtnListener()
            } else {
                fragContainer.innerHTML = loginInnerHtml
                isRegisterView = false
                setSubmitBtnListener()
            }
            tabs[k].classList.remove("on")
        }
    })
})
setSubmitBtnListener()

function setSubmitBtnListener() {
    // 必须重新获取，因为重写了页面
    const submitBtn = document.querySelector("#submit-btn")
    const getVerificationCode = document.querySelector("#get-verification-code")
    submitBtn.addEventListener('click', () => {
        if (isRegisterView) {
            const inputPhone = document.querySelector("#input-phone")
            const inputVerificationCode = document.querySelector("#input-verification-code")
            register("+86" + inputPhone.value, inputVerificationCode.value, "sms")
        } else {
            const inputId = document.querySelector("#input-id")
            const inputPassword = document.querySelector("#input-password")
            login(inputId.value, inputPassword.value, "password")
        }
    })
    getVerificationCode.addEventListener('click', async () => {
        if (isRegisterView) {
            // 发送验证码
            const inputPhone = document.querySelector("#input-phone")
            const res = await getVerifyCode("sms", "%2B86" + inputPhone.value)
            switch (res.status) {
                case 20001: {
                    alert('验证码已发送!')
                    break
                }
                default: {
                    alert('输入的手机号不正确')
                }
            }
        } else {
            // 忘记密码
            dialogForgetPass.style.display = "block"
        }
    })
}

async function login(account, token, type) {
    const formData = new FormData()
    formData.append("account", account)
    formData.append("token", token)
    formData.append("type", type)
    const res = await fetch(BASE_URL + "/users/login", {
        method: "POST",
        body: formData,
    })
    const obj = await res.json()
    switch (obj.status) {
        case 20000: {
            // 成功辣
            localStorage.setItem("access_token", obj.data.access_token)
            localStorage.setItem("refresh_token", obj.data.refresh_token)
            alert('登录成功')
            window.location.href = '../index.html'
            break
        }
        default: {
            alert("登录失败")
        }
    }
}

async function register(account, token, type) {
    const formData = new FormData()
    formData.append("account", account)
    formData.append("token", token)
    formData.append("type", type)
    const res = await fetch(BASE_URL + "/users/register", {
        method: "POST",
        body: formData,
    })
    const obj = await res.json()
    switch (obj.status) {
        case 20000: {
            localStorage.setItem("access_token", obj.data.access_token)
            localStorage.setItem("refresh_token", obj.data.refresh_token)
            alert('注册成功')
            window.location.href = '../index.html'
            break
        }
        default: {
            alert("注册失败")
        }
    }
}
