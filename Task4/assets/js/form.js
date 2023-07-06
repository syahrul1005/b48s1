function submitData(event) {
    event.preventDefault()

let name = document.getElementById("name").value
let email = document.getElementById("email").value
let phone = document.getElementById("phone-number").value
let subject = document.getElementById("subject").value
let message = document.getElementById("message").value

if (name == "" & email == "" & phone == "" & subject == "" & message == "") {
    return alert("Please complete all your Data")
} else if (name == "") {
    return alert("Please type your Name")
} else if (email == "") {
    return alert("Please type your Email")
} else if (phone == "") {
    return alert("Please type your Phone Number")
} else if (subject == "") {
    return alert("Please choose according to you purpose")
} else if (message == "") {
    return alert("Please type your Message")
}

let allData = {
    name,
    email,
    phone,
    subject,
    message,
}

console.log(allData)

const emailTo = "syahrulwafa67@gmail.com"

let a = document.createElement('a')
a.href = `mailto:${emailTo}?subject=${subject}&body=Hallo my name ${name}%0A%0A${message}.%0AThanks you,%0A%0APlease contact me: ${phone}`
a.click()
}