var frmSignupEl = document.getElementById("frm-signup")
var inpUsernameEl = document.getElementById("inp-username")
var inpPasswordEl = document.getElementById("inp-password")
var inpEmailEl = document.getElementById("inp-email")

var signupStatusEl = document.getElementById("signup-status")
var signupStatusMessageEl = document.getElementById("signup-status-message")


// Handle user registration request
frmSignupEl.addEventListener("submit", function(e) {

    // Suppress default form submission
    e.preventDefault()
    signupStatusEl.style.visibility = "visible"

    // Get user inputs
    // TODO: Trim space before injection to LOGIN class
    var username = inpUsernameEl.value.trim().toLowerCase()
    var email = inpEmailEl.value.trim().toLowerCase()
    var password = inpPasswordEl.value.trim()

    console.log(username)
    console.log(password)
    console.log(email)

    // Register user
    var registerClass = new Login()
    var responseSignup = registerClass.signup(username, password, email)

    console.log(responseSignup)

    // Inform user about signup status
    if (responseSignup === 200) {
        signupStatusMessageEl.innerHTML = "You're successfully signed up!"
        signupStatusMessageEl.setAttribute("class", "text-success")
        signupStatusEl.appendChild(signupStatusMessageEl)
    } 
    else if (responseSignup === 99) {
        signupStatusMessageEl.innerHTML = "Username is taken. Please try another username!"
        signupStatusMessageEl.setAttribute("class", "text-danger")
        signupStatusEl.appendChild(signupStatusMessageEl)
    } 
    else {
        signupStatusMessageEl.innerHTML = "Registration failed for unknown reason! Please try again now or later!"
        signupStatusMessageEl.setAttribute("class", "text-danger")
        signupStatusEl.appendChild(signupStatusMessageEl)
    }

    // Clear form fields
    frmSignupEl.reset()
})

// Clear status message fields
inpUsernameEl.addEventListener("input", function(e) {

    // Hide message container
    signupStatusEl.style.visibility = "hidden"
})