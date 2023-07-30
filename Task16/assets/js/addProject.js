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
        return alert("Please complete your Project Name")
    } 
    else if (startDate == "") {
        return alert("Please complete your Start Date")
    } 
    else if (endDate == "") {
        return alert("Please complete your End Date")
    } 
    else if (description == "") {
        return alert("Please complete your Description")
    } 
    else if (image == "") {
        return alert("Please upload your Image")
    } 
    else if (technologies == "") {
        return alert("Please choose your Technologies")
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
    console.log(dataBlog);

    renderBlog();
}

function renderBlog() {
    document.getElementById("previewProject").innerHTML = " "

    for(let x = 0; x < dataBlog.length; x++) {

        document.getElementById("previewProject").innerHTML += `
        <div class="card shadow bg-body p-2" style="width: 18rem;">
                    <a href="project-detail.html">
                        <img src="${dataBlog[x].image}" class="card-img-top object-fit-cover" style="max-height: 180px;"></a>
                        <div class="card-body card-body d-flex flex-column">
                        <h5 class="card-title m-0">${dataBlog[x].title}</h5>
                        <p>durasi: ${dataBlog[x].durasion}</p>
                        <p class="card-text">
                            ${dataBlog[x].description}
                        </p>
                        <div class="d-flex flex-column mt-auto">
                            <div class="fs-3">
                                ${dataBlog[x].technologies}
                            </div>
                            <div class="d-flex justify-content-between mt-4">
                                <button type="button" class="btn btn-dark rounded btnEdit">edit</button>
                                <button type="button" class="btn btn-dark rounded btnDelete">delete</button>
                            </div>
                        </div>
                    </div>
                </div>`
    }
}