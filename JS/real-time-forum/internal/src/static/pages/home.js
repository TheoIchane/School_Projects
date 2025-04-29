import { displayUsers } from "./users.js"

let page = document.getElementById('page')

export const displayHome = async () => {
    let resp = await fetch(`/api/home`)
    let r = await resp.json()
    document.title = r.title
    let posts_array = r.data.Posts
    console.log(posts_array)
    let posts = ""
    posts_array.forEach(post => {
        posts += formatPost(post) + " \n"
    });
    let content = ""
    if (r.user) {
        document.querySelector('.login').setAttribute('style', 'display:none')
        document.querySelector('.logged').setAttribute('style', 'display:flex')
        content = `
    <div class="page_title">
    Welcome ${r.user.username} !
    </div>
    <div id="posts">
    ${posts}
    </div>
    `
        displayUsers()
    } else {
        document.querySelector('.logged').setAttribute('style', 'display:none')
        document.querySelector('.login').setAttribute('style', 'display:flex')
        content = `
        <div class="title">
            Login or Register to create or interact a post.
        </div>
    `
        document.querySelector('#users_box').style.display = 'none'
    }
    page.innerHTML = content
}

function formatPost(post) {
    return `
    <div class="post">
        <p class="title">${post.title} </p>
        <div class="not_title">
        <div class="info">
            Author:
            <span class="username" >${post.author.username}</span>
            ${post.date} <br>
            Topic:
            ${post.topic.title}
        </div>
        <div class="content">${post.content} </div>
        </div>
    </div>
    `
}