// privat class # hanya bisa dipanggil di classnya
// di class variable = properties dan function = methods



class Testimonial {
    //enkapsulasi
    #image = ""
    #quote = ""
    
    constructor(image, quote) {
        this.#image = image
        this.#quote = quote
    }
    //abstraction
    get image() {
        return this.#image
    }

    get quote() {
        return this.#quote
    }

    get author() {
        throw new Error('there is must be author to make testimonial')
    }

    get testimonialHTML() {
        return `<div class="testimonial">
        <img src="${this.image}" alt="">
        <p class="quotes">"${this.quote}"</p>
        <p class="author"><b>- ${this.author}</b></p>
    </div>`
    }
}

//inheritance class child yang mendapat warisan dari parents (super).  extends
class UserTestimonial extends Testimonial {
    #user = ""

    constructor(image, quote, user) {
        super(image, quote)
        this.#user = user
    }
    //polymorphism
    get author() {
        return 'user : ' + this.#user
    }
}

class CompanyTestimonial extends Testimonial {
    #company = ""

    constructor(image, quote, company) {
        super(image, quote)
        this.#company = company
    }

    get author() {
        return "company : " + this.#company
    }
}

//class = object
const testimonial1 = new UserTestimonial("https://images.unsplash.com/photo-1512544783971-fb9a0691eda5?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1374&q=80", "I am Iron Man", "Tony Stark")
const testimonial2 = new UserTestimonial("https://images.unsplash.com/photo-1491328480217-db46cf0876c5?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=872&q=80", "I Don't Want To Go", "Peter Parker")
const testimonial3 = new CompanyTestimonial("https://images.unsplash.com/photo-1517993037474-692208825419?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=869&q=80", "Changing The World For A Better Future", "Stark Industries")


let testimonialAll = [testimonial1, testimonial2, testimonial3]

let testimonialHTML = ""

//x += y setara dengan x = x + y
for (let i = 0; i < testimonialAll.length; i++) {
    testimonialHTML += testimonialAll[i].testimonialHTML
}

document.getElementById("container-content").innerHTML = testimonialHTML