import { displayHome } from "./home.js"

export const getPost = async (id = 0) => {
    let r
    if (id == 0) {
        let resp = await fetch(`/api/post/`, {
            method: "GET"
        })
        r = await resp.json()
    } else {
        let resp = await fetch(`/api/post/${id}`)
        r = await resp.json()
    }
    return r
}

export const makePostPage = async () => {
    document.title = "Make a Post"
    let topics = await getTopics()
    let content = `
    <div class="post_form">
            <h2>Make a Post!</h2>
            <div class="body">
            <label for="title">Title:</label>
            <input id="title" type="text">
            <select name="topic" id="topic_id">
            <option disabled selected value="">Select an option</option>
            ${formatTopics(topics)}
            </select>
            <label for="content"></label>
            <input id="content" type="text-area">
            <button id="make_post">Create</button>
            </div>
            <div class="bottom">
            </div>
    </div>
        `
    page.innerHTML = content
    document.getElementById('make_post').addEventListener('click', postRequest)
}
async function postRequest() {
    let title = document.getElementById("title").value
    let content = document.getElementById("content").value
    let topic_id = Number(document.getElementById("topic_id").value)
    let body = JSON.stringify({
        title: title,
        topic_id: topic_id,
        content: content
    })
    let resp = await fetch("/api/post/", {
        method: "POST",
        body: body
    })
    let r = await resp.json()
    if (!r.error) {
        displayHome()
    } else {

    }
}

function formatTopics(topics) {
    let content = ""
    for (let i = 0; i < topics.length; i++) {
        let topic = topics[i]
        let opt = `<option value="${topic.id}">${topic.title}</option>`
        content += `${opt}`
    }
    return content
}

export const getTopics = async () => {
    let resp = await fetch("/api/topics/")
    topics = await resp.json()
}