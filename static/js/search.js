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
        data.data.forEach(item => {
            const result = document.createElement("search-result")
            result.setAttribute("id", item.mid)
            result.setAttribute("title", item.name)
            const scoreObj = JSON.parse(item.score)
            const detail = JSON.parse(item.detail)
            result.setAttribute("score", scoreObj.score)
            result.setAttribute("total", scoreObj.total_cnt)
            result.setAttribute("type", detail.type.join(","))
            result.setAttribute("authors", detail.director + "," + [...detail.writers,...detail.characters].join(","))
        })
        break
    }
    default: {
        alert("获取数据失败: " + data.info)
    }
}