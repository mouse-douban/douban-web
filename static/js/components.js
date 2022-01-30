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
            margin: 0 25px 10px 0;
        }

        img {
            display: block;
            float: left;
        }

        p {
            font-size: 13px;
            color: #37a;
            max-lines: 2;
            overflow: hidden;
        }

        span {
            color: orange;
        }

        #bottom {
            align-self: center;
            width: fit-content;
            height: 44px;
        }

        </style>
        <img width="115px" height="170px">
        <div id="bottom">
            <p></p>
        </div>
        `
        const shadow = this.attachShadow({ mode: 'open' })
        // insert
        shadow.appendChild(template.content.cloneNode(true))
        // attributes
        const src = this.getAttribute("src")
        const movie = this.getAttribute("movie")
        const score = this.getAttribute("score")
        // 获取元素 & 赋予属性
        const img = shadow.querySelector("img")
        img.src = src
        shadow.querySelector('p').innerHTML = `
            ${movie} <span>${score}</span>
        `
    }
}

customElements.define("movie-card", MovieCard)