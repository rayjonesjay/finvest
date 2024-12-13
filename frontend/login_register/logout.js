// Get elements
var logoutEl = document.getElementById("logout")

// Handle logout request
logoutEl.addEventListener("click", function(e) {

    // Clear login status
    localStorage.setItem("loginStatus", false)
    
    // Redirect user to login page
    location.assign("index.html")
})