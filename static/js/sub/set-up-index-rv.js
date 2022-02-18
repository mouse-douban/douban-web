import { getSubjects } from "../api.js"
import { getAbsolutePath } from "../utils.js"

const rv = new RecyclerView()
rv.setAdapter((data, index) => {
    const card = document.createElement("movie-card")
    card.setAttribute("src", data.avatar)
    card.setAttribute("movie", data.name)
    card.setAttribute("score", data.score)
    card.addEventListener("click", () => {
        localStorage.setItem("movieId", data.mid)
        location.href = getAbsolutePath("/static/movie")
    })
    card.flush()
    return card
})
rv.style = "position: relative;left: 225px;width: fit-content;margin-top: 20px;width: 650px;"
document.querySelector("body").appendChild(rv);
(async function () {
    const data = await getSubjects()
    if (data.status == 20000) {
        rv.setData(data.data)
    }
    rv.flush()
})()