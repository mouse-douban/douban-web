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