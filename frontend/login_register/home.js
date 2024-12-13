// Get elements
var homeGreetEl = document.getElementById("home-greet")
var greetUserEl = document.getElementById("greet-user")

// Greet if a user loggedin successfully
if (localStorage.getItem("loginStatus")) {

    var loginStatus = JSON.parse(localStorage.getItem("loginStatus"))

    if (loginStatus) {
        var userInfo = location.search
        var username = userInfo.split("=")[1]

        greetUserEl.textContent = username.toUpperCase().concat("!")
        homeGreetEl.appendChild(greetUserEl)
    } else {

        // Redirect user to login page
        location.assign("index.html")
    }
} else {
    // Redirect user to login page
    location.assign("index.html")
}