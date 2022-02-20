import { getCelebrityInfo } from "./api.js";
import { setup } from "./top-bar-status.js";

setup()

const id = localStorage.getItem("celebrityId")
// const id = 1026319

const data = await getCelebrityInfo(id)

switch (data.status) {
    case 20000: {
        document.title = `${data.data.name} - ${data.data.name_en} `
        document.querySelector("#content h1").textContent = data.data.name + " " + data.data.name_en
        for (let k in data.data) {
            const ele = document.querySelector(`#${k}`)
            if (k === "avatar") {
                ele.src = data.data[k]
                continue
            }
            if (ele === null) continue
            ele.textContent = data.data[k]
        }
        break
    }
    default: {
        alert("获取数据失败: " + data.info)
    }
}