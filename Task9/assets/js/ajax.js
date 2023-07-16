const promise = new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest()

    xhr.open("GET", "https://api.npoint.io/c9cc2e998ba1bb9784ee", true)
    xhr.onload = function () {
        // http code : 200 -> OK
        if (xhr.status === 200) {
            resolve(JSON.parse(xhr.responseText))
        } else if (xhr.status >= 400) {
            reject("Error loading data")
        }
    }
    xhr.onerror = function () {
        reject("Network error")
    }
    xhr.send()
})


let testimonialData = []

async function getData(rating) {
    try {
        const response = await promise
        console.log(response)
        testimonialData = response
        allTestimonial()
    } catch (err) {
        console.log(er)
    }
}

getData()




function allTestimonial() {
    let testimonialHTML = ""

    testimonialData.forEach((card,index) => {
        testimonialHTML += `<div class="testimonial">
        <img src="${card.image}" alt="">
        <p class="quotes">"${card.quote}"</p>
        <p class="author"><b>- ${card.user}</b></p>
        <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
    </div>`
    })

    document.getElementById("container-content").innerHTML = testimonialHTML
}

allTestimonial()

function filterTestimonial(rating) {
    let filteredTestimonialHTML = ""

    const filteredData = testimonialData.filter((card, ) => {
        return card.rating === rating //x===y >> x = x==y
    })

    filteredData.forEach((card) => {
        filteredTestimonialHTML += `<div class="testimonial">
        <img src="${card.image}" alt="">
        <p class="quotes">"${card.quote}"</p>
        <p class="author"><b>- ${card.user}</b></p>
        <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
    </div>`
    })

    document.getElementById("container-content").innerHTML = filteredTestimonialHTML
}