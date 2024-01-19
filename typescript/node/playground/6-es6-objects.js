
// Object property shorthand

const name = 'Oliver'
const userAge = '35'

const user = {
    name,
    age: userAge,
    location: 'Bogota'
}

console.log(user)

// Object Destructuring

const product = {
    label: 'Red Notebook',
    price: 3,
    stock: 20,
    salePrice: undefined
}

// const label = product.label
// const stock = product.stock

// const {label:productLabel, stock, rating = 5} = product
// console.log(productLabel)
// console.log(stock)
// console.log(rating)


const transaction = (type, { label, stock }) => {
    console.log(label, stock)

}

transaction('order', product)
