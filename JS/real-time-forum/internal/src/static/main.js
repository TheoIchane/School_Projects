import { displayHome } from "./pages/home.js";
import { getPost, getTopics, makePostPage } from "./pages/post.js"
import { loginPage, registerPage, logout } from "./pages/login.js";

document.getElementById('home').addEventListener('click', displayHome)

document.getElementById('login').addEventListener("click", loginPage)

document.getElementById('register').addEventListener('click', registerPage)

document.getElementById('logout').addEventListener('click', logout)

document.getElementById('post').addEventListener('click', makePostPage)

document.getElementById('mode').addEventListener('click', () => {
    let rev = {
        light: "dark",
        dark: "light"
    }
    document.body.id = rev[document.body.id]
})

// document.getElementById('allposts').addEventListener('click', getPost())

// document.getElementById('post1').addEventListener('click', getPost(1))

document.addEventListener('DOMContentLoaded', () => {
    displayHome()
})