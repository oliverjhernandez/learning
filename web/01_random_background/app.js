const but = document.querySelector("#changeColor");
const body = document.querySelector("body");
const h1 = document.querySelector("h1");

function changeColor() {
  new_color = pickNewColor();
  body.style.backgroundColor = new_color;
  h1.innerText = new_color;
}

function pickNewColor() {
  const red = Math.floor(Math.random(0, 255) * 255);
  const blue = Math.floor(Math.random(0, 255) * 255);
  const green = Math.floor(Math.random(0, 255) * 255);
  return `rgb(${red}, ${blue}, ${green})`;
}

but.addEventListener("click", changeColor);
