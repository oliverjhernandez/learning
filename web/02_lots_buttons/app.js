function pickNewColor() {
  const red = Math.floor(Math.random(0, 255) * 255);
  const blue = Math.floor(Math.random(0, 255) * 255);
  const green = Math.floor(Math.random(0, 255) * 255);
  return `rgb(${red}, ${blue}, ${green})`;
}

function main() {
  new_color = pickNewColor();
  this.style.backgroundColor = new_color;
}

const buttons = document.querySelectorAll("button");

for (const but of buttons) {
  but.addEventListener("click", main);
}
