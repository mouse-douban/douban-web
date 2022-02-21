import { addComment, getCelebrityInfo, getSubjectById, addReviews } from "./api.js";
import { setup } from "./top-bar-status.js";

setup()

const data = await getSubjectById(localStorage.getItem("movieId"))

const mainpic = document.querySelector("#mainpic")
const title = document.querySelector(".movie-title")
const subtitles = document.querySelectorAll(".sub-title")
const introduce = document.querySelector("#introduce")
const actorContainer = document.querySelector("#actor-container")

const dialogClosers = document.querySelectorAll(".dialog-closer")
const dialogBoxes = document.querySelectorAll(".dialog-box")
const dialogButtons = document.querySelectorAll(".action-line span")
const dialogSubmitBtns = document.querySelectorAll(".dialog-button")

dialogClosers.forEach(closer => {
    closer.addEventListener("click", () => {
         dialogBoxes.forEach(box => {
             box.style.display = "none"
         })
    })
})

dialogButtons.forEach((button, index) => {
    button.addEventListener("click", () => {
        dialogBoxes[index].style.display = "block"
    })
})

dialogSubmitBtns.forEach((button, index) => {
    button.addEventListener("click", () => {
        const box = dialogBoxes[index]
        switch (index){
            case 0: {
                // step 1 数据获取&校验
                // state 为 0 时，表示没有尚未评分
                const score = box.querySelector("ranking-stars").state
                if (score === 0) {
                    alert("请先评分")
                    break
                }
                const tag = box.querySelector(`input[type="text"]`).value
                if (tag.length === 0) {
                    alert("请输入短评标签")
                    break
                }
                const comment = box.querySelector("textarea").value
                if (comment.length === 0) {
                    alert("请输入短评内容")
                    break
                }
                const type = box.querySelector(`input[type="radio"]`).checked ? "before" : "after"
                // step 2 数据提交
                addShortComment(type, score, tag, comment)
                dialogBoxes[index].style.display = "none"
                break
            }
            case 1: {
                const score = box.querySelector("ranking-stars").state
                if (score === 0) {
                    alert("请先评分")
                    break
                }
                const tag = box.querySelector(`input[type="text"]`).value
                if (tag.length === 0) {
                    alert("请输入影评标题")
                    break
                }
                const comment = box.querySelector("textarea").value
                if (comment.length === 0) {
                    alert("请输入影评内容")
                    break
                }
                addReview(score, tag, comment)
                dialogBoxes[index].style.display = "none"
                break
            }
        }
    })
})

switch(data.status) {
    case 20000: {
        const movie = data.data
        mainpic.src = movie.avatar
        title.innerHTML = `${movie.name} <span style="color: #888888;font-size: 25px;">(${movie.date.substring(0, 4)})</span>`
        document.title = `${movie.name} (豆瓣)`
        subtitles.forEach(subtitle => {
            subtitle.textContent = subtitle.textContent.replace("电影名称", movie.name)
        })
        introduce.textContent = movie.plot
        subtitles[1].textContent = subtitles[1].textContent.replace("?", movie.celebrities.length)
        // actorContainer.innerHTML = ""
        movie.celebrities.forEach(async celebrity => {
            const data = await getCelebrityInfo(celebrity)
            const card = document.createElement("actor-card")
            card.setAttribute("src", data.data.avatar)
            card.setAttribute("actor", data.data.name)
            card.setAttribute("job", data.data.job)
            card.setAttribute("id", data.data.id)
            actorContainer.appendChild(card)
        })
        for (let k in movie.detail) {
            if (k === "website") continue
            const v = movie.detail[k]
            const item = document.querySelector("#" + k)
            if (Array.isArray(v)) {
                item.textContent = v.join(",")
            } else {
                if (k == "release") {
                    item.textContent = v.replace("00:00:00", "")
                    continue
                }
                item.textContent = v
            }
        }
        const bars = document.querySelectorAll(".bar")
        const spans = document.querySelectorAll(".bar-value")
        bars[0].style.width = movie.score.five.replace("%", "") + "px"
        spans[0].textContent = movie.score.five
        bars[1].style.width = movie.score.four.replace("%", "") + "px"
        spans[1].textContent = movie.score.four
        bars[2].style.width = movie.score.three.replace("%", "") + "px"
        spans[2].textContent = movie.score.three
        bars[3].style.width = movie.score.two.replace("%", "") + "px"
        spans[3].textContent = movie.score.two
        bars[4].style.width = movie.score.one.replace("%", "") + "px"
        spans[4].textContent = movie.score.one
        document.querySelector("#score").textContent = movie.score.score
        const starBox = document.querySelectorAll("#evaluate-box-top h2 span")[1]
        starBox.innerHTML = ""
        const emptyStarUrl = "https://img3.doubanio.com/f/shire/95cc2fa733221bb8edd28ad56a7145a5ad33383e/pics/rating_icons/star_hollow_hover@2x.png"
        const fullStarUrl = "https://img3.doubanio.com/f/shire/7258904022439076d57303c3b06ad195bf1dc41a/pics/rating_icons/star_onmouseover@2x.png"
        let count = 5
        for (let i = 1; i < movie.score.score / 2; i++) {
            starBox.innerHTML += `<img src="${fullStarUrl}" width="16" height="16">`
            count--
        }
        for (let i = 0; i < count; i++) {
            starBox.innerHTML += `<img src="${emptyStarUrl}" width="16" height="16">`
        }
        break
    }
    default: {
        alert(data.info)
    }
}

async function addShortComment(type, score, tag, comment) {
    const data = await addComment(localStorage.getItem("movieId"), score, type, comment, tag)
    switch (data.status) {
        case 20220: {
            alert("成功发布短评")
            location.reload()
            break
        }
        default: {
            alert(data.info)
        }
    }
}

async function addReview(score, name, content) {
    const data = await addReviews(localStorage.getItem("movieId"), score, name, content)
    switch (data.status) {
        case 20220: {
            alert("成功发布影评")
            location.reload()
            break
        }
        default: {
            alert(data.info)
        }
    }
}