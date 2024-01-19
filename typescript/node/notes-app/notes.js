const fs = require('fs');
const chalk = require('chalk');

const getNotes = () => {
    return "Your notes..."
}

const loadNotes = () => {
    try {
        const dataBuffer = fs.readFileSync('notes.json');
        const dataJSON = dataBuffer.toString();
        return JSON.parse(dataJSON);
    } catch (e) {
        return []
    }
}

const writeNotes = (array) => {
    const dataJSON = JSON.stringify(array);
    fs.writeFileSync('notes.json', dataJSON); 
}

const addNotes = (title, body) => {
    const notes = loadNotes();

    const duplicateNote = notes.find((note) => note.title === title )

    if (!duplicateNote) {
        notes.push({
            title: title,
            body: body
        })
        console.log(notes);
        writeNotes(notes)
    } else {
        console.log('Title already taken');
    }

}

const removeNote = (title) => {
    const notes = loadNotes();

    const filteredNotes = notes.filter((note) => note.title != title )

    if (notes.length > filteredNotes.length) {
        console.log(chalk.blue('Note removed!')) ;
        writeNotes(filteredNotes);
    } else {
       console.log(chalk.red('No note found!')) ;
    }

}

const listNotes = () => {
    const notes = loadNotes();
    console.log(chalk.inverse.blue('List of Notes:'))
    notes.forEach((note, index) => {
        console.log(`${index + 1}. ${chalk.blue(note.title)}`);
    })
}

const readNote = (title) => {
    const notes = loadNotes();
    const readNote = notes.find((note) => note.title === title)

    if (readNote) {
        console.log("Title: ", chalk.inverse.green(readNote.title));
        console.log("Body:  ", readNote.body);
    } else {
        console.log(chalk.red("Note not found!"))
    }
}


module.exports = {
    getNotes: getNotes,
    addNotes: addNotes,
    removeNote: removeNote,
    listNotes: listNotes,
    readNote: readNote
}

