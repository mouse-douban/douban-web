import { getSubjects } from "./api.js";
import { setup } from "./top-bar-status.js";

setup()

const radios = document.querySelectorAll("input[type=radio]")

radios.forEach(radio => {
    radio.addEventListener("change", async () => {
        if (radio.checked) {
            const sort = radio.id == "sort-latest" ? "" : "hotest"
            const tabLayout = document.querySelector("tab-layout")
            const rv = document.querySelector("recycler-view")
            const tag = tabLayout.seleted ==  "热门" ? "" : tabLayout.seleted
            const res = await getSubjects(0, 20, sort, tag)
            rv.setData(res.data)
            rv.flush()
        }
    })
})
