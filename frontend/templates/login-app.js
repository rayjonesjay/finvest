// Get DOM elements
var frmSigninEl = document.getElementById("frm-signin")

var inpUsernameEl = document.getElementById("inp-username")
var inpPasswordEl = document.getElementById("inp-password")
var signinStatusEl = document.getElementById("signin-status")
var signinStatusMessageEl = document.getElementById("signin-status-message")

var btnForgetPasswordEl = document.getElementById("btn-forgetPassword")
var btnSignupEl = document.getElementById("btn-signup")

const showPasswordEl = document.getElementById("show-password")


// Handle signin request
frmSigninEl.addEventListener("submit", function(e) {

    // Suppress default form submission
    e.preventDefault()
    signinStatusEl.style.visibility = "visible"

    // Create login class
    // TODO: Trim space before injection to LOGIN class
    var username = inpUsernameEl.value.trim().toLowerCase()
    var password = inpPasswordEl.value.trim()

    console.log(username)
    console.log(password)

    // Allow user to signin
    var signinClass = new Login()
    var responseSignin = signinClass.signin(username, password)

    // Inform user about registration
    if (responseSignin === 200) {

        // Get url specific to a user
        // url_home = "home.html?username=".concat(username)
        url_home = "home.html?username=".concat(username)

        // Redirect user to home page
        location.assign(url_home)
    } else if (responseSignin === 99) {
        signinStatusMessageEl.innerHTML = "Login failed due to wrong username or password! Please try again."
        signinStatusMessageEl.setAttribute("class", "text-danger")
        signinStatusEl.appendChild(signinStatusMessageEl)
    } else if (responseSignin === 9) {
        signinStatusMessageEl.innerHTML = "Registraion failed becuase you're not registed yet! Please register first."
        signinStatusMessageEl.setAttribute("class", "text-danger")
        signinStatusEl.appendChild(signinStatusMessageEl)
    } else {
        signinStatusMessageEl.innerHTML = "Registration failed for unknown reason! Please try again now or later!"
        signinStatusMessageEl.setAttribute("class", "text-danger")
        signinStatusEl.appendChild(signinStatusMessageEl)
    }
    
    // Clear form fields
    frmSigninEl.reset()
})

// Remove signin status container
inpUsernameEl.addEventListener("input", function (e) {

    // Clear signin status message container
    // TODO: For some reason, old signin status message (wrong password) is shown just before redirection to index.html
    signinStatusEl.style.visibility = "hidden"
})

// Handle show password request
showPasswordEl.addEventListener("change", function (e) {

    // Show password
    if (e.target.checked) {
        // Change input field into a TEXT type
        inpPasswordEl.type = "text"
    } else {
        inpPasswordEl.type = "password"
    }
})

// Handle forget password request
btnForgetPasswordEl.addEventListener("click", function(e) {
    // Redirect user to forget password page
    location.assign("forget-password.html")
})

// Handle signup request
btnSignupEl.addEventListener("click", function(e) {

    // Redirect user to signup page
    location.assign("signup.html")
})