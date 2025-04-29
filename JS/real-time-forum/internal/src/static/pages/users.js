import { displayChatBox } from "./chat.js"
let users_box = document.getElementById('users_box')

export const displayUsers = async () => {
    users_box.style.display = 'block'
    users_box.innerHTML = "<h2>USERS</h2>"
    let resp = await fetch("/api/users")
    let res = await resp.json()
    if (res.error) return
    console.log(res)
    let online = res.online ? sortUserArray(res.online) : []
    let offline = res.offline ? sortUserArray(res.offline) : []
    online.forEach((x) => createDOM(x, "online"))
    offline.forEach((x) => createDOM(x, "offline"))
}

function sortUserArray(array = []) {
    return array.sort((a, b) => {
        if (a.username.toLowerCase() > b.username.toLowerCase()) return 1
        if (a.username.toLowerCase() == b.username.toLowerCase()) return 0
        if (a.username.toLowerCase() < b.username.toLowerCase()) return -1
    })
}

function createDOM(user, statut) {
    let button = document.createElement('div')
    button.classList.add("user", statut)
    button.innerHTML = `
    <span>.</span> ${user.username} (${user.lastname} ${user.firstname})
    `
    button.onclick = () => {
        displayChatBox(user.username)
    }
    users_box.appendChild(button)
}