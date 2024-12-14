

class Login {

    constructor () {
        
        // Initialize validation trigger values
        this.validationSignupTrigger = 0   
        this.validationSigninTrigger = 0  

        this.validationForgetPasswordTriggerObject = {
            triggerStatusValue: 0,
            username: '',
            password: ''
        } 
    }

    // A method to signin a user
    signin = function (username, password) {

        // Check whether any user is registered in the localStorage
        if (localStorage.getItem("users")) {

            // Get existing users
            var users = JSON.parse(localStorage.getItem("users"))

            // Get all usernames
            var usernameAll = users.map(a => a.username)

            if (usernameAll.includes(username)) {
                
                // Perform signin authentication
                // TODO [optional] - Possible to extract a user and check the password, instead of looping...
                for (let i = 0; i < users.length; i++) {

                    if (users[i]["username"] === username) {
                        if (users[i]["password"] === password) {
                            this.validationSigninTrigger = 200  // Login is successful
                        } else {
                            this.validationSigninTrigger = 99  // The value <99> indicates that user exists, but password doesn't match!
                        }

                    }
                } 
            } else {
                this.validationSigninTrigger = 9  // The value <9> indicates that the user is unknown!
            }
        } else {
            this.validationSigninTrigger = 9  // The value <9> indicates that the localStorage is empty. Thus, the user is not registred yet!
        }

        // Update login status - needed to perfom a logouy operation!
        if (this.validationSigninTrigger === 200) {
            localStorage.setItem("loginStatus", true)
        } else {
            localStorage.setItem("loginStatus", false)
        }

        return this.validationSigninTrigger  
    }

    // A method to register a new user
    signup = function (username, password, email) {

        // Add user to database
        var newUser = {
            'id': uuidv4(),
            'username': username,
            'password': password,
            'email': email,
        }

        // Check whether a username is taken
        var isUsernameTaken = this.checkWhetherUsernameTaken(newUser)

        // Add new user to existing users list
        if (isUsernameTaken) {
            this.validationSignupTrigger = 99   // The value <99> indicates that username is taken
        } else {
            // Proceed registering the new user
            if (localStorage.getItem("users")) {
                var users = JSON.parse(localStorage.getItem("users"))
            } else {
                // Handle a case where no user registered before
                var users = []
            }

            // Add the new user to the localStorage
            users.push(newUser)
            localStorage.setItem("users", JSON.stringify(users))

            // Update response status code
            this.validationSignupTrigger = 200   // Response code <200> indicates that registration went successful!
        }

        return this.validationSignupTrigger
    }

    // A method to check whether a username is taken
    checkWhetherUsernameTaken = function (newUserObject) {

        // Check whether username is taken
        if (localStorage.getItem("users")) {

            // Get existing users
            var users = JSON.parse(localStorage.getItem("users"))

            // Check whether a new username matches with existing usernames
            for (let i = 0; i < users.length; i++) {
                if (users[i]["username"] === newUserObject["username"]) {
                    return true
                }
            }

            return false
        } else {
            return false
        }
    }

    // A method to help user find username and password
    forgetPassword = function (userEmail) {

        // Check whether any user is registered in the localStorage
        if (localStorage.getItem("users")) {

            // Get existing users
            var users = JSON.parse(localStorage.getItem("users"))

            // Get all emails
            var emailAll = users.map(a => a.email)

            if (emailAll.includes(userEmail)) {

                // Grab username and password
                // TODO [optional] - The looping might be optimized...
                for (let i = 0; i < users.length; i++) {

                    if (users[i]["email"] === userEmail) {

                        const userFound = users[i]

                        this.validationForgetPasswordTriggerObject["triggerStatusValue"] = 200
                        this.validationForgetPasswordTriggerObject["username"] = userFound["username"]
                        this.validationForgetPasswordTriggerObject["password"] = userFound["password"]
                    }
                }
            } else {
                this.validationForgetPasswordTriggerObject["triggerStatusValue"] = 9  // The value <9> indicates that the email is unknown!
            }
        } else {
            this.validationForgetPasswordTriggerObject["triggerStatusValue"] = 9  // The value <9> indicates that the localStorage is empty. Thus, the user is not registred yet!
        }

        return this.validationForgetPasswordTriggerObject
    }
}