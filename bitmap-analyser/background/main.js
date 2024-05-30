import './style.css'

document.querySelector('#app').innerHTML = `
<div class="container">
    <input type="file" id="file">
    <img src="https://placehold.it/700x400">
</div>
`

const input = document.querySelector("#file")
const img = document.querySelector("img")
const container = document.querySelector(".container")
input.addEventListener("change", async (event) => {
  const file = event.target.files[0]
  const url = URL.createObjectURL(file)
  img.src = url
  console.log(file)
  const formData = new FormData()
  formData.append("file", file) 
  const response = await fetch("http://localhost:9090/background", {
    method: "POST",
    mode: "cors",
    contentType: "multipart/form-data",
    body: formData
  })
  const color = await response.text()
  console.log("response :", color)
  container.style.backgroundColor = color
})

