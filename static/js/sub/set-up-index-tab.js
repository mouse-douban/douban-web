import { getSubjects } from "../api.js"

const types = [
    "热门",
    "喜剧",
    "爱情",
    "动作",
    "科幻",
    "悬疑",
    "惊悚",
    "动画",
    "奇幻"
]

let recyclerView = null

const tl = new TabLayout(types, async tab => {
    if (recyclerView == null) {
        recyclerView = document.querySelector("recycler-view")
    }
    const tag = tab == "热门" ? "" : tab
    const res = await getSubjects(0, 20, "latest", tag)
    if (res.status == 20000) {
        recyclerView.setData(res.data)
        recyclerView.flush()
    }
    return true
})
tl.style = "position: relative;left: 225px;width: fit-content;"
document.querySelector("body").appendChild(tl)