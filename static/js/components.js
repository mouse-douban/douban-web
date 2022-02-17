// 组件化
class MovieCard extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
        :host {
            width: fit-content;
            flex-direction: column;
            display: flex;
            justify-content: center;
            margin: 0 25px 10px 0;
        }

        img {
            margin-bottom: 5px;
        }

        p {
            font-size: 13px;
            color: #37a;
            max-lines: 2;
            overflow: hidden;
            text-align: center;
        }

        span {
            color: orange;
        }

        .card {
            box-shadow: 0 1px 1px 0 rgba(0,0,0,0.2), 0 1px 1px 0 rgba(0,0,0,0.19);
        }

        #card {
            cursor: pointer;
            border-radius: 10px;
            height: fit-content;
            width: fit-content;
            box-sizing: border-box;
            padding: 0px 10px 0px 10px;
            display: flex;
            flex-direction: column;
        }
        </style>
        <div id="card" class= "card">
            <img width="115px" height="170px">
            <p></p>
        </div>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        // insert
        this.shadow.appendChild(template.content.cloneNode(true))
        // attributes
        const src = this.getAttribute("src")
        const movie = this.getAttribute("movie")
        const score = this.getAttribute("score")
        // 获取元素 & 赋予属性
        const img = this.shadow.querySelector("img")
        img.src = src
        this.shadow.querySelector('p').innerHTML = `
            ${movie} <span>${score}</span>
        `
    }

    // 属性改变回调
    attributeChangedCallback(name, oldValue, newValue) {
        if (name === "src") {
            this.shadow.querySelector("img").src = newValue
        }
        if (name === "movie") {
            this.shadow.querySelector('p').innerHTML = `
                ${newValue} <span>${this.getAttribute("score")}</span>
            `
        }
        if (name === "score") {
            this.shadow.querySelector('p').innerHTML = `
                ${this.getAttribute("movie")} <span>${newValue}</span>
            `
        }
    }

    flush() {
        this.shadow.querySelector("img").src = this.getAttribute("src")
        this.shadow.querySelector('p').innerHTML = `
            ${this.getAttribute("movie")} <span>${this.getAttribute("score")}</span>
        `
    }
}

customElements.define("movie-card", MovieCard)

class UserReview extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
            #review {
                margin: 10px;
                display: inline-flex;
                height: fit-content;
                min-height: 100px;
                width: 450px;
                padding: 10px;
                box-sizing: border-box;
                border-radius: 10px;
                cursor: pointer;
            }
            #stars {
                width: fit-content;
                padding: 20px;
                height: fit-content;
                color: #ffc83d;
            }
            #content-box {
                width: 100%;
                min-height: 100px;
            }
            #content {
                min-height: 80%;
                font-size: 16px;
                margin-bottom: 10px;
                font-style: italic;
                padding-top: 22px;
                padding-left: 10px;
                color: #9a9c9a;
                box-sizing: border-box;
            }
            #movie-name {
                width: 100%;
                height: fit-content;
                align-self: flex-end;
                justify-self: right;
                font-size: 13px;
                color: #37a;
                text-align: right;
                font-style: italic;
                
            }
            .card {
                box-shadow: 0 1px 1px 0 rgba(0,0,0,0.2), 0 1px 1px 0 rgba(0,0,0,0.19);
            }
        </style>
        <div id="review" class="card">
            <div id="stars">${this.getAttribute("score")}⭐</div>
            <div id="content-box">
                <div id="content">“ ${this.getAttribute("content")} ”</div>
                <div id="movie-name">《${this.getAttribute("movie")}》</div>
            </div>
        </div>
        `
        const shadow = this.attachShadow({ mode: 'open' })
        shadow.appendChild(template.content.cloneNode(true))
    }

}

customElements.define("user-review", UserReview)

class UserMovieList extends HTMLElement {

    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
        #list {
            padding: 20px;
            display: flex;
            width: fit-content;
            max-width: 100%;
            flex-wrap: wrap;
            height: fit-content;
            box-sizing: border-box;
            border-radius: 10px;
            justify-content: center;
            align-items: flex-start;
            min-height: 150px;
            min-width: 300px;
            margin-bottom: 20px;
        }
        movie-card {
            margin: 10px;
        }
        .card {
            box-shadow: 0 1px 1px 0 rgba(0,0,0,0.2), 0 1px 1px 0 rgba(0,0,0,0.19);
        }
        #name {
            font-family: 'Noto Sans', sans-serif;
        }
        </style>
        <h2 id="name">${this.getAttribute("name")}</h2>
        <div id="list" class="card">
        </div>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        this.shadow.appendChild(template.content.cloneNode(true))
        this.list = this.shadow.querySelector("#list")
        JSON.parse(this.getAttribute("data")).forEach(movieInfo => {
            const movieCard = new MovieCard()
            movieCard.setAttribute("src", movieInfo.avatar)
            movieCard.setAttribute("movie", movieInfo.movie)
            movieCard.setAttribute("score", movieInfo.score.score)
            this.list.appendChild(movieCard)
            movieCard.flush()
        })
    }

    attributeChangedCallback(name, oldValue, newValue) {
        if (name === "name") {
            this.shadow.querySelector("#name").innerHTML = newValue
        }
        if (name === "data") {
            this.list.innerHTML = ""
            JSON.parse(newValue).forEach(movieInfo => {
                const movieCard = document.createElement("movie-card")
                movieCard.setAttribute("src", movieInfo.avatar)
                movieCard.setAttribute("movie", movieInfo.movie)
                movieCard.setAttribute("score", movieInfo.score.score)
                this.list.appendChild(movieCard)
            })
        }
    }
}

customElements.define("user-movie-list", UserMovieList)

// 类似Android的TabLayout组件化实现
// 使用前需要设置属性 tabs 和 onTabItemClick
// 也许之后可以试试实现ViewPager2
// 果然还得是组件化才彳亍啊
class TabLayout extends HTMLElement {
    /**
     * 
     * @param {*} tabs tab名称数组
     * @param {*} onTabItemClick tabItem的点击回调
     */
    constructor(tabs, onTabItemClick) {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
            :host {
                display: flex;
                flex-direction: row;
                flex-wrap: wrap;
                align-items: flex-start;
                justify-content: flex-start;
                height: fit-content;
                margin: 10px;
            }
            .tab {
                padding: 5px 10px;
                margin: 5px;
                font-size: 16px;
                color: #666666;
                border-radius: 3px;
                transition: all 0.3s;
                cursor: pointer;
            }
            .selected {
                background-color: #4b8ccb;
                color: #ffffff;
            }
            .tab:hover {
                background-color: #eeeeee;
            }
        </style>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        this.shadow.appendChild(template.content.cloneNode(true))
        this.tabs = tabs
        this.onTabItemClick = onTabItemClick
        this.tabItems = []
        this.selected = tabs[0]
        this.tabs.forEach(tab => {
            const tabItem = document.createElement("div")
            tabItem.classList.add("tab")
            tabItem.textContent = tab
            tabItem.addEventListener("click", () => {
                if (tabItem.classList.contains("selected")) {
                    return
                }
                if (this.onTabItemClick(tab)) {
                    this.selected = tab
                    this.tabItems.forEach(item => item.classList.remove("selected"))
                    tabItem.classList.add("selected")
                }
            })
            this.shadow.appendChild(tabItem)
            this.tabItems.push(tabItem)
        })
        this.tabItems[0].classList.add("selected")
    }
}

customElements.define("tab-layout", TabLayout)

class SeparatorLine extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <div style="position: relative;left: 230px;width: 650px;height: 10px;border-bottom: 1px solid #eeeeee;"></div>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        this.shadow.appendChild(template.content.cloneNode(true))
    }
}

customElements.define("separator-line", SeparatorLine)

// 有时间也许可以尝试一下实现懒加载
class RecyclerView extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
            :host {
                display: flex;
                flex-warp: wrap;
                justify-content: flex-start;
                align-items: flex-start;
            }

            .element {
                margin: 10px;
            }

        </style>
        <div id="list">
        </div>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        this.shadow.appendChild(template.content.cloneNode(true))
        this.list = this.shadow.querySelector("#list")
    }

    // 更新数据
    flush() {
        this.list.innerHTML = ""
        this.datas.forEach((data, index) => {
            const ele = this.adapter(data, index)
            this.list.appendChild(ele)
        })
    }

    /**
     * 设置数据
     * 
     * @param {array} datas
     */
    setData(datas) {
        this.datas = datas
    }

    /**
     * 添加单条数据
     * 
     * @param {*} data 
     */
    appendData(data) {
        this.datas.push(data)
    }

    /**
     * 给这个recyclerview设置适配回调函数
     * 
     * @param {*} adapter (data: any, index: number) => HTMLElement
     */
    setAdapter(adapter) {
        this.adapter = adapter
    }
}

customElements.define("recycler-view", RecyclerView)

class ActorCard extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
        :host {
            width: fit-content;
            flex-direction: column;
            display: flex;
            justify-content: center;
            margin: 0 25px 10px 0;
        }

        img {
            margin-bottom: 5px;
        }

        p {
            font-size: 14px;
            color: #37a;
            max-lines: 2;
            margin: 5px;
            overflow: hidden;
            text-align: center;
            transition: all 0.3s;
        }

        .actor:hover {
            color: white;
            background-color: #3377aa;
        }

        .job {
            color: #494949;
        }

        .card {
            box-shadow: 0 1px 1px 0 rgba(0,0,0,0.2), 0 1px 1px 0 rgba(0,0,0,0.19);
        }

        #card {
            cursor: pointer;
            border-radius: 10px;
            height: fit-content;
            width: fit-content;
            box-sizing: border-box;
            padding: 0px 10px 0px 10px;
            display: flex;
            flex-direction: column;
        }
        </style>
        <div id="card" class= "card">
            <img width="115px" height="170px">
            <p class="actor"></p>
            <p class="job"></p>
        </div>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        // insert
        const node = template.content.cloneNode(true)
        this.shadow.appendChild(node)
    }

    connectedCallback() {
        this.shadow.querySelector('.actor').innerHTML = this.getAttribute("actor")
        this.shadow.querySelector('.job').innerHTML = this.getAttribute("job")
        this.shadow.querySelector('img').src = this.getAttribute("src")
    }

    // 属性改变回调
    attributeChangedCallback(name, oldValue, newValue) {
        if (name === "src") {
            this.shadow.querySelector("img").src = newValue
        }
        if (name === "actor") {
            this.shadow.querySelector('.actor').innerHTML = `${newValue}`
        }
        if (name === "job") {
            this.shadow.querySelector('.job').innerHTML = `${this.getAttribute("job")}`
        }
    }
}

customElements.define("actor-card", ActorCard)

class ShortComment extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
        :host {
            width: 100%;
        }

        span {
            font-size: 13px;
        }

        .score {
            color: orange;
        }

        .time {
            color: #aaaaaa;
        }

        .clickable {
            color: #3377aa;
            cursor: pointer;
            transition: all 0.3s;
        }

        .clickable:hover {
            color: white;
            background-color: #3377aa;
        }

        .top-box {            
            border-top: 1px solid #efefef;
            padding: 10px 0;
        }

        .content {
            color: #494949;
            font-size: 14px;
        }

        </style>
        <div class= "top-box">
            <span class="user clickable">user</span> <span class="type">看过</span> <span class="score-span"><span class="score">5</span>⭐</span> <span class="time">2022-02-01 10:46:17</span>
            <span style="float: right;"><span class="stars">5801</span> <span class="stars-text clickable">有用</span></span>
        </div>
        <p class="content">短评内容</p>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        // insert
        const node = template.content.cloneNode(true)
        this.shadow.appendChild(node)
    }

    connectedCallback() {
        this.shadow.querySelector(".type").innerHTML = this.getAttribute("type") === "after" ? "看过" : "想看"
        this.shadow.querySelector('.user').innerHTML = this.getAttribute("user")
        this.shadow.querySelector('.score').innerHTML = this.getAttribute("score")
        this.shadow.querySelector('.time').innerHTML = this.getAttribute("time")
        this.shadow.querySelector('.content').innerHTML = this.getAttribute("content")
        this.shadow.querySelector('.stars').innerHTML = this.getAttribute("stars")
    }

    // 属性改变回调
    attributeChangedCallback(name, oldValue, newValue) {
        if (name === "user") {
            this.shadow.querySelector('.user').innerHTML = `${newValue}`
        }
        if (name === "score") {
            this.shadow.querySelector('.score').innerHTML = `${newValue}`
        }
        if (name === "time") {
            this.shadow.querySelector('.time').innerHTML = `${newValue}`
        }
        if (name === "content") {
            this.shadow.querySelector('.content').innerHTML = `${newValue}`
        }
        if (name === "stars") {
            this.shadow.querySelector('.stars').innerHTML = `${newValue}`
        }
        if (name === "type") {
            this.shadow.querySelector('.type').innerHTML = newValue === "after" ? "看过" : "想看"
        }
    }
}

customElements.define("short-comment", ShortComment)

class MovieComment extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
        :host {
            width: 100%;
        }

        span {
            font-size: 13px;
        }

        .score {
            color: orange;
        }

        .time {
            color: #aaaaaa;
        }

        .clickable {
            color: #3377aa;
            cursor: pointer;
            transition: all 0.3s;
        }

        .clickable:hover {
            color: white;
            background-color: #3377aa;
        }

        .top-box {            
            border-top: 1px solid #efefef;
            padding: 10px 0;
        }

        .content {
            color: #494949;
            font-size: 14px;
        }

        .title {
            margin: 0;
            font-weight: normal;
            font-size: 16px;
            width: fit-content;
        }

        </style>
        <div class= "top-box">
            <img src="" width="36px" height="36px"> <span class="user clickable">user</span> <span class="score-span"><span class="score">5</span>⭐</span> <span class="time">2022-02-01 10:46:17</span>
            <span style="float: right;"><span class="stars">5801</span> <span class="stars-text clickable">有用</span></span>
        </div>
        <h2 class="title clickable">影评标题</h2>
        <p class="content">影评内容</p>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        // insert
        const node = template.content.cloneNode(true)
        this.shadow.appendChild(node)
    }

    connectedCallback() {
        this.shadow.querySelector('.user').innerHTML = this.getAttribute("user")
        this.shadow.querySelector('.score').innerHTML = this.getAttribute("score")
        this.shadow.querySelector('.time').innerHTML = this.getAttribute("time")
        this.shadow.querySelector('.content').innerHTML = this.getAttribute("content")
        this.shadow.querySelector('.stars').innerHTML = this.getAttribute("stars")
        this.shadow.querySelector('.title').innerHTML = this.getAttribute("title")
        this.shadow.querySelector('img').src = this.getAttribute("user-icon")
    }

    attributeChangedCallback(name, oldValue, newValue) {
        if (name === "user") {
            this.shadow.querySelector('.user').innerHTML = `${newValue}`
        }
        if (name === "score") {
            this.shadow.querySelector('.score').innerHTML = `${newValue}`
        }
        if (name === "time") {
            this.shadow.querySelector('.time').innerHTML = `${newValue}`
        }
        if (name === "content") {
            this.shadow.querySelector('.content').innerHTML = `${newValue}`
        }
        if (name === "stars") {
            this.shadow.querySelector('.stars').innerHTML = `${newValue}`
        }
        if (name == "title") {
            this.shadow.querySelector('.title').innerHTML = `${newValue}`
        }
        if (name == "user-icon") {
            this.shadow.querySelector('img').src = `${newValue}`
        }
    }
}

customElements.define("movie-comment", MovieComment)