function sendPostRequest() {
    const textInput = document.getElementById("text-input");
    const text = textInput.value;
    const banner = document.querySelector('input[name="banner"]:checked').value;

    const formData = new FormData();
    formData.append("text", text);
    formData.append("banner", banner);

    fetch("/ascii-art", {
        method: "POST",
        body: formData
    })
    .then(response => {
        if (response.ok) {
            return response.text();
        } else {
            throw new Error("HTTP error: " + response.status);
        }
    })
    .then(data => {
        const resultDiv = document.getElementById("result");
        resultDiv.innerHTML = data;
    })
    .catch(error => {
        console.error(error);
    });
}
