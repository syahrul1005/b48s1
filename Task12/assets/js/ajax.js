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
        testimonialHTML += `<div class="card shadow bg-body p-2" style="width: 18rem;">
        <img src="${card.image}" class="card-img-top object-fit-cover pb-2"
            style="max-width: 270px; max-height: 190px;">
        <div class="card-body p-1">
            <p class="card-text text-justify fst-italic w-75 " style="height: 60px;">${card.quote}
            </p>
            <div class="text-end">
                <p class=" m-0">- ${card.user}</p>
                <p class=" m-0">${card.rating} <i class="fa-solid fa-star"></i></p>
            </div>
        </div>
    </div>`
    })

    document.getElementById("testimonial").innerHTML = testimonialHTML
}

allTestimonial()

function filterTestimonial(rating) {
    let filteredTestimonialHTML = ""

    const filteredData = testimonialData.filter((card, ) => {
        return card.rating === rating //x===y >> x = x==y
    })

    filteredData.forEach((card) => {
        filteredTestimonialHTML += `<div class="card shadow bg-body p-2" style="width: 18rem;">
        <img src="${card.image}" class="card-img-top object-fit-cover pb-2"
            style="max-width: 270px; max-height: 190px;">
        <div class="card-body p-1">
            <p class="card-text text-justify fst-italic w-75 " style="height: 60px;">${card.quote}
            </p>
            <div class="text-end">
                <p class=" m-0">- ${card.user}</p>
                <p class=" m-0">${card.rating} <i class="fa-solid fa-star"></i></p>
            </div>
        </div>
    </div>`
    })

    document.getElementById("testimonial").innerHTML = filteredTestimonialHTML
}