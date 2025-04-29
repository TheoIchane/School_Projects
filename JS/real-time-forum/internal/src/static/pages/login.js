let page = document.getElementById('page')
let messagePopUp = document.getElementById('messagePopUp')

export const loginPage = async () => {
    document.title = "Login"
    let content = `
<div class="log_form">
        <h2>Login</h2>
        <div class="body">
        <label for="login_idtf">Identifier:</label>
        <input id="login_idtf" type="text">
        <label for="login_pswd">Password:</label>
        <input id="login_pswd" type="text">
        <button id="login_req">Login</button>
        </div>
        <div class="bottom">
        Not registered yet ? <br>  
        <button id="reg">Register</button>
        </div>
</div>
    `
    page.innerHTML = content
    document.getElementById('login_req').addEventListener('click', loginRequest)
    document.getElementById('reg').addEventListener('click', registerPage)
}

export const registerPage = async () => {
    document.title = "Register"
    let content = `
<div class="log_form">
        <h2>Register</h2>
        <div class="body">
            <label for="reg_email">Email:</label>
            <input id="reg_email" type="text" required>
            <label for="reg_pswd">Password:</label>
            <input id="reg_pswd" type="password" required>
            <label for="reg_username">Username:</label>
            <input id="reg_username" type="text" required>
            <label for="reg_lastname">LastName:</label>
            <input id="reg_lastname" type="text" required>
            <label for="reg_firstname">FirstName:</label>
            <input id="reg_firstname" type="text" required>
            <label for="reg_age>">Age</label>
            <input id="reg_age" type="text" required>
            <label for="reg_gender">Gender:</label>
            <select name="gender" id="reg_gender" required>
                <option disabled selected value="">-- Select an option --</option>
                <option value="Male">Male</option>
                <option value="Female">Female</option>
                <option value="Non Binary">Non Binary</option>
                <option value="Other">Other</option>
                <option value="Prefer not to answer">Prefer not to answer</option>
            </select>
            <button id="register_req">Register</button>
        </div>
        <div class="bottom">
        Already registered ? <br>  
        <button id="alr_reg">Login</button>
        </div>
</div>
    `
    page.innerHTML = content
    document.getElementById('register_req').addEventListener('click', registerRequest)
    document.getElementById('alr_reg').addEventListener('click', loginPage)
}

export const logout = async () => {
    let resp = await fetch("/api/logout")
    let r = await resp.json()
    displayHome()
}



async function loginRequest() {
    let identifier = document.getElementById('login_idtf').value
    let password = document.getElementById('login_pswd').value
    let body = JSON.stringify({
        identifier: identifier,
        password: password
    })
    let resp = await fetch("/api/login", {
        method: "POST",
        body: body
    })
    let r = await resp.json()
    if (r.status == 200) {
        displayHome()
    } else {

    }
}

async function registerRequest() {
    let email = document.getElementById('reg_email').value
    let password = document.getElementById('reg_pswd').value
    let lastname = document.getElementById('reg_lastname').value
    let firstname = document.getElementById('reg_firstname').value
    let age = document.getElementById('reg_age').value
    let gender = document.getElementById('reg_gender').value
    let username = document.getElementById('reg_username').value
    // console.log(email, password, lastname, firstname, age, gender, username)
    let body = JSON.stringify({
        email: email,
        username: username,
        firstname: firstname,
        lastname: lastname,
        gender: gender,
        age: Number(age),
        password: password
    })
    let resp = await fetch("/api/register", {
        method: "POST",
        body: body
    })
    let r = await resp.json()
    if (r.message) {
        messagePopUp.innerText = r.message
        setTimeout(() => {
            messagePopUp.innerText = ""
        }, 5000)
        if (!r.error) {
            displayHome()
        }
    }
}