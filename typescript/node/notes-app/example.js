

const fs = require('fs');
const add = require('./utils')
const getNotes = require('./notes')

const sum = add(3, 4);
console.log(sum);

const myNotes = getNotes();
console.log(myNotes);

const filename = 'notes.txt'
fs.writeFileSync(filename, "Hello madafaka\n")

fs.appendFileSync(filename, "Hellow biatch\n")
