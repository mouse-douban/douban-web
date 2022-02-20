import { search } from "./api.js";
import { setup } from "./top-bar-status.js";

setup()

const container = document.querySelector("#search-result-container")
const data = await search(localStorage.getItem("searchKey"))

document.querySelector("h1").textContent = `搜索 ${localStorage.getItem("searchKey")}`
document.querySelector("#inp-query").value = localStorage.getItem("searchKey")
container.innerHTML = ""
switch (data.status) {
    case 20000: {
        data.data[1].forEach(item => {
            const result = document.createElement("search-result")
            result.setAttribute("id", item.mid)
            result.setAttribute("title", item.name)
            result.setAttribute("score", item.score.score)
            result.setAttribute("total", item.score.total_cnt)
            result.setAttribute("types", item.detail.type.join(","))
            result.setAttribute("authors", item.detail.director + "," + [...item.detail.writers,...item.detail.characters].join(","))
            result.setAttribute("avatar", item.avatar)
            container.appendChild(result)
        })
        break
    }
    default: {
        alert("获取数据失败: " + data.info)
    }
}