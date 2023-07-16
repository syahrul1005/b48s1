let navBotOpen = false
function openNavBot () {
    let navBottom = document.getElementById("navBot")
    if (!navBotOpen) {
        navBottom.style.display = "block"
        navBotOpen = true
    } else {
        navBottom.style.display = "none"
        navBotOpen = false
        
    }
}