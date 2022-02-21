// 组件化
class MovieCard extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
        :host {
            width: fit-content;
            display: inline-block;
            margin: 0 25px 10px 0;
            transition: all 0.3s;
        }

        img {
            margin: 10px 0px;
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
            height: 250px;
            width: 135px;
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
    }

    connectedCallback() {
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
                min-height: 50%;
                margin-top: 15px;
                font-size: 16px;
                font-style: italic;
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
            h3 {
                margin: 0;
            }
            .card {
                box-shadow: 0 1px 1px 0 rgba(0,0,0,0.2), 0 1px 1px 0 rgba(0,0,0,0.19);
            }
        </style>
        <div id="review" class="card">
            <div id="stars">${this.getAttribute("score")}⭐</div>
            <div id="content-box">
                <h3 id="title">${this.getAttribute("title")}</h3>
                <div id="content">“ ${this.getAttribute("content")} ”</div>
                <div id="movie-name">《${this.getAttribute("movie")}》</div>
            </div>
        </div>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        this.shadow.appendChild(template.content.cloneNode(true))
    }

    connectedCallback() {
        this.shadow.querySelector("#stars").textContent = this.getAttribute("score") + "⭐"
        this.shadow.querySelector("#content").textContent = `"${this.getAttribute("content")}"`
        this.shadow.querySelector("#movie-name").textContent = `《${this.getAttribute("movie")}》`
        this.shadow.querySelector("#title").textContent = this.getAttribute("title")
    }

    attributeChangedCallback(name, oldValue, newValue) {
        if (name === "score") {
            this.shadow.querySelector("#stars").textContent = newValue + "⭐"
        }
        if (name === "content") {
            this.shadow.querySelector("#content").textContent = newValue
        }
        if (name === "movie") {
            this.shadow.querySelector("#movie-name").textContent = newValue
        }
        if (name === "title") {
            this.shadow.querySelector("#title").textContent = newValue
        }
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
            width: 100%;
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
        <h2 id="name"></h2>
        <div id="list" class="card">
        </div>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        this.shadow.appendChild(template.content.cloneNode(true))
        this.list = this.shadow.querySelector("#list")
    }

    connectedCallback() {
        JSON.parse(this.getAttribute("data")).forEach(movieInfo => {
            const movieCard = new MovieCard()
            movieCard.setAttribute("src", movieInfo.avatar)
            movieCard.setAttribute("movie", movieInfo.movie)
            movieCard.setAttribute("score", movieInfo.score.score)
            this.list.appendChild(movieCard)
            movieCard.flush()
        })
        this.shadow.querySelector("#name").textContent = this.getAttribute("name")
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
            #list {
                display: inline-flex;
                flex-wrap: wrap;
                width: 60%;
                height: fit-content;
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
            width: 135px;
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
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
            margin: 5px;
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
            width: 135px;
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
        // 获取绝对地址
        function getAbsolutePath(path) {
            const curWwwPath = window.document.location.href
            const pathName = window.document.location.pathname;
            const pos = curWwwPath.indexOf(pathName)
            const localhostPaht = curWwwPath.substring(0, pos)
            return localhostPaht + path
        }
        this.shadow.querySelector('.actor').innerHTML = this.getAttribute("actor")
        this.shadow.querySelector('.job').innerHTML = this.getAttribute("job")
        this.shadow.querySelector('img').src = this.getAttribute("src")
        this.shadow.querySelector('#card').addEventListener("click", () => {
            localStorage.setItem("celebrityId", this.getAttribute("id"))
            window.open(getAbsolutePath("/static/celebrities"))
        })
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

class RecyclerTail extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
        :host {
            width: 100%;
        }

        .tail {
            margin-bottom: 20px;
            height: fit-content;
            padding: 20px;
            background-color: #efefef;
            border-top: 1px solid #efefef;
            display: flex;
            justify-content: center;
            align-items: center;
            cursor: pointer;
            transition: all 0.3s;
            border-radius: 5px;
        }

        .tail:hover {
            background-color: #f2f2f2;
        }

        .tail-text {
            font-size: 16px;
            color: #aaaaaa;
        }

        </style>
        <div class="tail">
            <span class="tail-text">点此加载更多</span>
        </div>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        const node = template.content.cloneNode(true)
        this.shadow.appendChild(node)
        const tail = this.shadow.querySelector(".tail")
        // state: hasMore noMore loading
        this.state = "hasMore"
        tail.addEventListener("click", async () => {
            if (this.state === "hasMore") {
                this.state = "loading"
                tail.textContent = `正在加载...`
                // 这里必须执行完毕才能切换状态回去，所以onloadmore必须是一个async的函数
                // onloadmore需要返回一个boolean来标注加载状态 true为加载完成 false为没有更多了
                const flag = await eval(this.onloadmore)
                if (flag) {
                    this.state = "hasMore"
                    tail.textContent = `点此加载更多`
                } else {
                    this.state = "noMore"
                    tail.textContent = `没有更多了`
                }
            }
        })
    }

    connectedCallback() {
        this.onloadmore = this.getAttribute("onloadmore")
    }
}

customElements.define("recycler-tail", RecyclerTail)

const emptyStarUrl = "https://img3.doubanio.com/f/shire/95cc2fa733221bb8edd28ad56a7145a5ad33383e/pics/rating_icons/star_hollow_hover@2x.png"
const fullStarUrl = "https://img3.doubanio.com/f/shire/7258904022439076d57303c3b06ad195bf1dc41a/pics/rating_icons/star_onmouseover@2x.png"
const keys = ["", "很差", "较差", "还行", "推荐", "力荐"]

// 评分组件
class RankingStars extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
            :host {
                padding: 10px;
            }

            img {
                cursor: pointer;
            }

            span {
                color: #888888;
                font-size: 15px;
            }
        </style>
        <img src="${emptyStarUrl}" id="1" width="16" height="16">
        <img src="${emptyStarUrl}" id="2" width="16" height="16">
        <img src="${emptyStarUrl}" id="3" width="16" height="16">
        <img src="${emptyStarUrl}" id="4" width="16" height="16">
        <img src="${emptyStarUrl}" id="5" width="16" height="16">
        <span></span>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        const node = template.content.cloneNode(true)
        this.shadow.appendChild(node)
        this.span = this.shadow.querySelector("span")
        this.stars = this.shadow.querySelectorAll("img")
        this.state = 0
        this.stars.forEach((star, index) => {
            star.addEventListener("mouseover", () => {
                this.stars.forEach(star => {
                    star.src = emptyStarUrl
                })
                for (let i = 0; i < star.id; i++) {
                    this.stars[i].src = fullStarUrl
                }
                this.span.textContent = keys[star.id]
            })
            star.addEventListener("mouseout", () => {
                this.stars.forEach(star => {
                    star.src = emptyStarUrl
                })
                for (let i = 0; i < this.state; i++) {
                    this.stars[i].src = fullStarUrl
                }
                this.span.textContent = keys[this.state]
            })
            star.addEventListener("click", () => {
                this.state = index + 1
            })
        })
    }
}

customElements.define("ranking-stars", RankingStars)

class UnMutableRankingStars extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
            span {
                color: #888888;
                font-size: 15px;
            }
        </style>
        <img src="${emptyStarUrl}" width="10" height="10">
        <img src="${emptyStarUrl}" width="10" height="10">
        <img src="${emptyStarUrl}" width="10" height="10">
        <img src="${emptyStarUrl}" width="10" height="10">
        <img src="${emptyStarUrl}" width="10" height="10">
        <span></span>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        const node = template.content.cloneNode(true)
        this.shadow.appendChild(node)
        this.span = this.shadow.querySelector("span")
        this.stars = this.shadow.querySelectorAll("img")
        this.state = 0
    }

    connectedCallback() {
        const score = this.getAttribute("score")
        this.shadow.querySelector("span").textContent = score
        for (let k = 0; k < score / 2 - 1; k++) {
            this.stars[k].src = fullStarUrl
        }
    }

    attributeChangedCallback(name, oldValue, newValue) {
        this.shadow.querySelector("span").textContent = newValue
        for (let k = 0; k < newValue / 2 - 1; k++) {
            this.stars[k].src = fullStarUrl
        }
    }
}

customElements.define("unmutable-ranking-stars", UnMutableRankingStars)

class SearchResult extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
        :host {
            display: flex;
            align-items: flex-start;
            justify-content: flex-start;
        }

        img {
            max-width: 100px;
            max-height: 100px;
            margin-right: 20px;
            margin-bottom: 20px;
        }

        #types,#authors {
            max-lines: 1;
            overflow: hidden;
            font-size: 13px;
            color: #888888;
        }

        h3 {
            margin: 0 0 10px 0;
            font-size: 16px;
            color: #3377aa;
            transition: all 0.3s;
            cursor: pointer;
            width: fit-content;
            font-weight: normal;
        }

        h3:hover {
            color: white;
            background-color: #3377aa;
        }

        div {
            color: #494949;
            font-size: 13px;
        }

        </style>
        <img src= "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2246432125.webp">
        <div>
            <h3>凉宫春日的消失 涼宮ハルヒの消失 (2010)</h3>
            <div style="margin-bottom: 10px;"><unmutable-ranking-stars score="9.1"></unmutable-ranking-stars> (<span>31722</span>人评价)</div>
            <div id="types">日本 / 动画 / 喜剧 / 科幻 / The Disappearance of Haruhi Suzumiya / Suzumiya Haruhi no shôshitsu / 162分钟</div>
            <div id="authors">石原立也 / 武本康弘 / 平野绫 / 杉田智和 / 茅原实里 / 青木沙耶香 / 后藤邑子 / 桑谷夏子 / 松元惠 / 松冈由贵</div>
        </div>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        const node = template.content.cloneNode(true)
        this.shadow.appendChild(node)
    }

    connectedCallback() {
        // 获取绝对地址
        function getAbsolutePath(path) {
            const curWwwPath = window.document.location.href
            const pathName = window.document.location.pathname;
            const pos = curWwwPath.indexOf(pathName)
            const localhostPaht = curWwwPath.substring(0, pos)
            return localhostPaht + path
        }
        const h3 = this.shadow.querySelector("h3")
        h3.addEventListener("click", () => {
            localStorage.setItem("movieId", this.getAttribute("id"))
            window.open(getAbsolutePath("/static/movie"))
        })
        h3.textContent = this.getAttribute("title")
        this.shadow.querySelector("unmutable-ranking-stars").setAttribute("score", this.getAttribute("score"))
        this.shadow.querySelector("#types").textContent = this.getAttribute("types")
        this.shadow.querySelector("#authors").textContent = this.getAttribute("authors")
        this.shadow.querySelector("div span").textContent = this.getAttribute("total")
        this.shadow.querySelector("img").src = this.getAttribute("avatar")
    }
}

customElements.define("search-result", SearchResult)

class SingleDiscussion extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>

        div {
            color: #494949;
            font-size: 16px;
            display: inline-block;
        }

        .container {
            border-top: 1px dashed #e6e6e6;
            padding: 10px 0;
        }

        .title {
            width: 400px
        }

        .line2 {
            width: 150px;
        }

        div span {
            color: #3377aa;
            transition: all 0.3s;
            cursor: pointer;
        }

        div span:hover {
            color: white;
            background-color: #3377aa;
        }
        </style>
        <div class="container">
            <div class="title"><span>标题</span></div>
            <div class="line2">来自<span class="author">寒雨</span></div>
            <div class="time">2022-02-20 14:54:35</div>
        </div>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        const node = template.content.cloneNode(true)
        this.shadow.appendChild(node)
    }

    connectedCallback() {
        // 获取绝对地址
        function getAbsolutePath(path) {
            const curWwwPath = window.document.location.href
            const pathName = window.document.location.pathname;
            const pos = curWwwPath.indexOf(pathName)
            const localhostPaht = curWwwPath.substring(0, pos)
            return localhostPaht + path
        }
        const title = this.shadow.querySelector(".title span")
        title.textContent = this.getAttribute("title")
        title.addEventListener("click", () => {
            localStorage.setItem("discussionId", this.getAttribute("id"))
            window.open(getAbsolutePath("/static/discussion"))
        })
        this.shadow.querySelector(".author").textContent = this.getAttribute("author")
        this.shadow.querySelector(".time").textContent = this.getAttribute("time")
    }
}

customElements.define("single-discussion", SingleDiscussion)

class DiscussionReply extends HTMLElement {
    constructor() {
        super()
        const template = document.createElement("template")
        template.innerHTML = `
        <style>
        
        .info {
            background-color: #f2fbf2;
            font-size: 15px;
        }

        .author {
            color: #3377aa;
            transition: all 0.3s;
            cursor: pointer;
        }

        .author:hover {
            color: white;
            background-color: #3377aa;
        }

        .bottom {
            margin-top: 10px;
            font-size: 15px;
            width: fit-content;
            float: right;
            color: #bbbbbb;
        }

        .bottom span {
            cursor: pointer;
            transition: all 0.3s;
        }

        .bottom span:hover {
            color: white;
            background-color: #bbbbbb;
        }

        .info {
            padding: 5px 3px;
        }

        </style>
        <div style="margin: 10px 0;">
            <img src="https://gitee.com/coldrain-moro/images_bed/raw/master/images/remilia.png" width="48px" height="48px">
            <div style="width: 92.5%;display: inline-block;margin-left: 20px;">
                <div class="info"><span style="color: #494949;">2021-11-02 17:02:09</span> <span class="author">寒雨</span></div>
                <p class="content">巴拉巴拉巴拉巴拉</p>
                <div class="bottom"> <span>赞</span> <span>回应</span></div>
            </div>
        </div>
        `
        this.shadow = this.attachShadow({ mode: 'open' })
        const node = template.content.cloneNode(true)
        this.shadow.appendChild(node)
    }

    connectedCallback() {
        const author = this.shadow.querySelector(".author")
        author.textContent = this.getAttribute("author")
        this.shadow.querySelector("img").src = this.getAttribute("avatar")
        this.shadow.querySelector(".info span").textContent = this.getAttribute("time")
        this.shadow.querySelector(".content").textContent = this.getAttribute("content")
    }

    attributeChangedCallback(name, oldValue, newValue) {
        if (name === "author") {
            this.shadow.querySelector(".author").textContent = newValue
        }
        if (name === "avatar") {
            this.shadow.querySelector("img").src = newValue
        }
        if (name === "time") {
            this.shadow.querySelector(".info span").textContent = newValue
        }
        if (name === "content") {
            this.shadow.querySelector(".content").textContent = newValue
        }
    }
}

customElements.define("discussion-reply", DiscussionReply)