import { getCelebrityInfo, getSubjectById } from "./api.js";
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

dialogClosers.forEach(closer => {
    closer.addEventListener("click", () => {
         dialogBoxes.forEach(box => {
             box.style.display = "none"
         })
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
        actorContainer.innerHTML = ""
        movie.celebrities.forEach(async celebrity => {
            const data = await getCelebrityInfo(celebrity).data
            const card = document.createElement("actor-card")
            card.setAttribute("src", data.avatar)
            card.setAttribute("actor", data.name)
            card.setAttribute("job", data.job)
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
        const spans = document.querySelectorAll("#evaluate-box-top p span")
        spans[0].textContent = movie.score.five
        spans[1].textContent = movie.score.four
        spans[2].textContent = movie.score.three
        spans[3].textContent = movie.score.two
        spans[4].textContent = movie.score.one
        document.querySelector("#score").textContent = movie.score.score
        const starBox = document.querySelectorAll("#evaluate-box-top h2 span")[1]
        starBox.textContent = ""
        for (let i = 1; i < movie.score.score / 2; i++) {
            starBox.textContent += "⭐"
        }
        break
    }
    default: {
        alert(data.info)
    }
}