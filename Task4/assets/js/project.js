let dataBlog = []

function addBlog(event) {
    event.preventDefault();

    let title = document.getElementById("name").value
    let startDate = new Date(document.getElementById("startDate").value)
    let endDate = new Date(document.getElementById("endDate").value)
    let description = document.getElementById("description").value
    let image = document.getElementById("image").files
    let icnnodejs = document.getElementById("nodejs").checked
    let icnreactjs = document.getElementById("reactjs").checked
    let icnvuejs = document.getElementById("vuejs").checked
    let icnjavacscript = document.getElementById("javascript").checked

    let interval = endDate - startDate //satuan milisecond 

    let intervalSeconds = Math.floor(interval / 1000) 
    let intervalMinutes = Math.floor(intervalSeconds / 60)
    let intervalHours = Math.floor(intervalMinutes / 60)
    let intervalDay = Math.floor(intervalHours / 24)
    let intervalMonth = Math.floor(intervalDay / 30)
    let intervalYear = Math.floor(intervalMonth / 12)

    let time = {
        intervalSeconds,
        intervalMinutes,
        intervalHours,
        intervalDay,
        intervalMonth,
        intervalYear,
    }

    let durasion = time

    if (intervalYear <= 0 & intervalMonth <= 0 & intervalDay <= 0 & intervalHours <= 0 & intervalMinutes <= 0 ) {
        durasion = intervalSeconds + " seconds "
    } 
    else if (intervalYear <= 0 & intervalMonth <= 0 & intervalDay <= 0 & intervalHours <= 0 ) {
        durasion = intervalMinutes + " minutes "
    } 
    else if (intervalYear <= 0 & intervalMonth <= 0 & intervalDay <= 0 ) {
        durasion = intervalHours + " hours "
    } 
    else if (intervalYear <= 0 & intervalMonth <= 0 ) {
        durasion = intervalDay + " day "
    } 
    else if (intervalYear <= 0 ) { 
        durasion = intervalMonth + " month " + (intervalDay-(intervalMonth*30)-1) + " day "
    } 
    else if (intervalYear >= 1 ) {
        durasion = intervalYear + " year " + (intervalMonth-(intervalYear*12)) + " month "
    }


    //kondisi icon ketika ceklist(true) maka tampil iconnya
    let tech = []

    if (icnnodejs) {
        tech.push('<i class="fa-brands fa-node-js"></i>')
    }
    if (icnreactjs) {
        tech.push('<i class="fa-brands fa-react"></i>')
    }
    if (icnvuejs) {
        tech.push('<i class="fa-brands fa-vuejs"></i>')
    }
    if (icnjavacscript) {
        tech.push('<i class="fa-brands fa-square-js"></i>')
    }

    let technologies = tech.join('  ')   //menggabungkan semua array ke dalam string

    //peringatan form belum terisi
    if (title == "" & startDate == "" & endDate == "" & description == "" & image == "" & technologies == "") {
        return alert("Please complete all your Data")
    } 
    else if (title == "") {
        return alert("Please complete all your Data")
    } 
    else if (startDate == "") {
        return alert("Please complete all your Data")
    } 
    else if (endDate == "") {
        return alert("Please complete all your Data")
    } 
    else if (description == "") {
        return alert("Please complete all your Data")
    } 
    else if (image == "") {
        return alert("Please complete all your Data")
    } 
    else if (technologies == "") {
        return alert("Please complete all your Data")
    }

    image = URL.createObjectURL(image[0])

    let project = {
        title,
        startDate,
        endDate,
        time,
        durasion,
        description,
        image,
        technologies
    }
    dataBlog.push(project);
    renderBlog();
    console.log(dataBlog);
}

function renderBlog() {
    document.getElementById("preview-project").innerHTML = ''

    for(let x = 0; x < dataBlog.length; x++) {

        document.getElementById("preview-project").innerHTML += `
        
            <div class="previewProject-container">
                <div class="preImage">
                    <a href="index-project-detail.html">
                        <img src="${dataBlog[x].image}">
                    </a>
                </div>
                <div class="preTitle">
                    <p>${dataBlog[x].title}</p>
                </div>
                <div class="preDuration">
                    <p>durasi : ${dataBlog[x].durasion}</p>
                </div>
                <div class="preDescription">
                    <p>${dataBlog[x].description}</p>
                </div>
                <div class="preTechnologies">
                    ${dataBlog[x].technologies}
                </div>
                <div class="button2">
                    <button><a href="index-project.html">edit</a></button>
                    <button>delete</button>
                </div>
            </div>
        `
    }
}