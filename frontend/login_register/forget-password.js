// Get DOM elements
var frmForgetPasswordEl = document.getElementById("frm-forget-password")
var emailEl = document.getElementById("inp-email")

var forgetPasswordStatusEl = document.getElementById("forget-password-status")
var forgetPasswordStatusMessageEl = document.getElementById("forget-password-status-message")

const btnSignupEl = document.getElementById("btn-signup")
const btnLoginEl = document.getElementById("btn-login")


// Handle forget password request
frmForgetPasswordEl.addEventListener("submit", function(e) {

    // Prevent default form submission
    e.preventDefault()
    forgetPasswordStatusEl.style.visibility = "visible"
    btnSignupEl.style.visibility = "visible"
    btnLoginEl.style.visibility = "visible"

    // Get user email
    var email = emailEl.value

    // Provide user a login info 
    var forgetPasswordClass = new Login()
    var responseForgetPasswordObject = forgetPasswordClass.forgetPassword(email)

    if (responseForgetPasswordObject["triggerStatusValue"] === 200) {

        // Update user about status
        forgetPasswordStatusMessageEl.innerHTML = `Username: ${responseForgetPasswordObject["username"]}` + "<br />" + `Password: ${responseForgetPasswordObject["password"]}`
        forgetPasswordStatusMessageEl.setAttribute("class", "text-success")
        forgetPasswordStatusEl.appendChild(forgetPasswordStatusMessageEl)

        // Redirect user to login page
        const btnLogin = document.createElement("a")
        btnLogin.setAttribute('href', "index.html")
        btnLogin.innerHTML = "Login"
        btnLogin.style.border = "3px solid"
        btnLogin.style.padding = "5px"
        btnLoginEl.appendChild(btnLogin)

        // Hide signup button
        btnSignupEl.style.visibility = "hidden"
    } else {

        // Update user about status
        forgetPasswordStatusMessageEl.innerHTML = `The email provided (${email}) is not registered yet. Please register!`
        forgetPasswordStatusMessageEl.setAttribute("class", "text-danger")
        forgetPasswordStatusEl.appendChild(forgetPasswordStatusMessageEl)

        // Display a register button to user
        var btnRegister = document.createElement("a")
        btnRegister.setAttribute('href', "index.html")
        btnRegister.innerHTML = "Signup"
        btnRegister.style.border = "3px solid"
        btnRegister.style.padding = "5px"
        btnSignupEl.appendChild(btnRegister)

        // Hide login button
        btnLoginEl.style.visibility = "hidden"
    }

    // Clear form input fields
    frmForgetPasswordEl.reset()
})

// Remove forgetPassword status container
emailEl.addEventListener("input", function (e) {

    // Hide status container and buttons
    forgetPasswordStatusEl.style.visibility = "hidden"
    btnSignupEl.style.visibility = "hidden"
    btnLoginEl.style.visibility = "hidden"
})