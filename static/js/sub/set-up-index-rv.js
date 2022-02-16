import { getSubjects } from "../api.js"

const rv = new RecyclerView()
rv.setAdapter((data, index) => {
    const card = document.createElement("movie-card")
    card.setAttribute("src", data.avatar)
    card.setAttribute("movie", data.name)
    card.setAttribute("score", data.score)
    card.flush()
    return card
})
rv.style = "position: relative;left: 225px;width: fit-content;margin-top: 20px;max-width: 600px;"
document.querySelector("body").appendChild(rv);
(async function () {
    const data = await getSubjects()
    if (data.status == 20000) {
        rv.setData(data.data)
    }
    rv.flush()
})()