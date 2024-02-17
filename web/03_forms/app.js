const form = document.querySelector("form");
const product = document.querySelector("#product");
const quantity = document.querySelector("#qty");
const list = document.querySelector("#list");

const addList = (event) => {
  event.preventDefault();
  const prod = product.value;
  const qty = quantity.value;
  const new_item = newListItem(`${qty} ${prod}`);
  list.appendChild(new_item);
  product.value = "";
  quantity.value = "";
};

const newListItem = (param) => {
  const new_item = document.createElement("li");
  const text = document.createTextNode(param);
  new_item.appendChild(text);
  return new_item;
};

form.addEventListener("submit", addList);
