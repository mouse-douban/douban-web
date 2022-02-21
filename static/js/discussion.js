import { getDiscussion, reply, starDiscussion } from "./api.js";
import { setup } from "./top-bar-status.js";

setup()

const id = localStorage.getItem("discussionId")
const data = await getDiscussion(id)
const submitBtn = document.querySelector(".submit")
const replyInput = document.querySelector(".reply-input")

switch(data.status) {
    case 20000: {
        const { name, username, avatar, stars, content, date } = data.data
        document.querySelector("#avatar").src = avatar
        document.querySelector(".user").textContent = username
        document.querySelector("#name").textContent= name
        document.querySelector("#stars").textContent = "赞 " + stars
        document.querySelector(".discussion-content").textContent = content
        document.querySelector(".date").textContent = date.replace("T", " ").replace("Z", "")
        break
    }
    default: {
        alert(data.data.detail)
    }
}

// 点赞
document.querySelector(".like-btn").addEventListener("click", async () => {
    const data = await starDiscussion(id)
    switch (data.status) {
        case 20220: {
            alert("点赞成功")
            location.reload()
            break
        }
        default: {
            alert(data.data.detail)
        }
    }
})

// 回应
submitBtn.addEventListener("click", async () => {
    const content = replyInput.value
    if (content.length === 0) {
        alert("回复内容不能为空")
        return
    }
    const data = await reply(id, "discussion", content)
    switch (data.status) {
        case 20000: {
            alert("回复成功")
            location.reload()
            break
        }
        default: {
            alert(data.data.detail)
        }
    }
})
