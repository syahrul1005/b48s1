// // privat class # hanya bisa dipanggil di classnya
// // di class variable = properties dan function = methods

// class Testimonial {
//     #image = ""
//     #quote = ""
    
//     constructor(image, quote) {
//         this.#image = image
//         this.#quote = quote
//     }

//     get image() {
//         return this.#image
//     }

//     get quote() {
//         return this.#quote
//     }

//     get author() {
//         throw new Error('there is must be author to make testimonial')
//     }

//     get testimonialHTML() {
//         return `<div class="testimonial">
//         <img src="${this.image}" alt="">
//         <p class="quotes">"${this.quote}"</p>
//         <p class="author"><b>- ${this.author}</b></p>
//     </div>`
//     }
// }

// class UserTestimonial extends Testimonial {
//     #user = ""

//     constructor(image, quote, user) {
//         super(image, quote)
//         this.#user = user
//     }

//     get author() {
//         return 'user : ' + this.#user
//     }
// }

// class CompanyTestimonial extends Testimonial {
//     #company = ""

//     constructor(image, quote, company) {
//         super(image, quote)
//         this.#company = company
//     }

//     get author() {
//         return "company : " + this.#company
//     }
// }

// const testimonial1 = new UserTestimonial("https://images.unsplash.com/photo-1512544783971-fb9a0691eda5?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1374&q=80", "I am Iron Man", "Tony Stark")
// const testimonial2 = new UserTestimonial("https://images.unsplash.com/photo-1491328480217-db46cf0876c5?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=872&q=80", "I Don't Want To Go", "Peter Parker")
// const testimonial3 = new CompanyTestimonial("https://images.unsplash.com/photo-1517993037474-692208825419?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=869&q=80", "Changing The World For A Better Future", "Stark Industries")


// let testimonialAll = [testimonial1, testimonial2, testimonial3]

// let testimonialHTML = ""

// for (let i = 0; i < testimonialAll.length; i++) {
//     testimonialHTML += testimonialAll[i].testimonialHTML
// }

// document.getElementById("container-content").innerHTML = testimonialHTML

const testimonialData = [
    {
        user: "Peter Parker",
        quote: "I Don't Want To Go",
        image: "https://images.unsplash.com/photo-1491328480217-db46cf0876c5?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=872&q=80",
        rating: 5
    },
    {
        user: "Tony Stark",
        quote: "I am... Iron Man",
        image: "https://images.unsplash.com/photo-1512544783971-fb9a0691eda5?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1374&q=80",
        rating: 4
    },
    {
        user: "Thanos",
        quote: "I am... inevitable",
        image: "https://images.unsplash.com/photo-1623944864235-db595bfccaad?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80",
        rating: 1
    },
    {
        user: "Stark Industries",
        quote: "Changing The World For A Better Future",
        image: "https://images.unsplash.com/photo-1517993037474-692208825419?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=869&q=80",
        rating: 4
    },
]



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