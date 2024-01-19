const fs = require('fs');

// const book = {
//     title: 'Ego is the enemy',
//     author: 'Ryan Holiday'
// }
// 
// const bookJSON = JSON.stringify(book);
// fs.writeFileSync('1-json.json', bookJSON);
// 
// const bookParse = JSON.parse(bookJSON);
// console.log(bookParse["author"]);
//

// const dataBuffer = fs.readFileSync('1-json.json');
// const dataJSON = dataBuffer.toString();
// const data = JSON.parse(dataJSON)
// console.log(data.title);


const sample = fs.readFileSync('samplejson.js');
const data = JSON.parse(sample);

data["name"] = "Oliver"
data["age"] = 35

console.log(data);

const dataJSON = JSON.stringify(data);

fs.writeFileSync('samplejson.js', dataJSON);
