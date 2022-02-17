

const rv = new RecyclerView()
rv.setAdapter((data, index) => {
    const card = document.createElement("movie-card")
    return card
})