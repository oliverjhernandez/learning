#!/usr/bin/env node

const validator = require('validator');
const chalk = require('chalk');
const yargs = require('yargs')
const notes = require('./notes');


////////////////////////////////////////////
//
// Yargs
//

yargs.version('1.2.0')

// Create add ommand
yargs.command({
    command: 'add',
    describe: 'Adds a new note',
    builder: {
        title: {
            describe: 'Note title.',
            demandOption: true,
            type: 'string'
        },
        body: {
            describe: 'Note body',
            demandOption: true,
            type: 'string'
        }
    },
    handler(argv) {
        notes.addNotes(argv.title, argv.body)
    }
})

// Create remove command
yargs.command({
    command: 'remove',
    describe: 'Removes a note',
    builder: {
        title: {
            describe: 'Note Title',
            demandOption: true,
            type: 'string'
        }
    },
    handler(argv) {
        notes.removeNote(argv.title)
    }
})

// Create remove command
yargs.command({
    command: 'list',
    describe: 'Lists all notes',
    handler() {
        notes.listNotes();
    }
})

// Create read command
yargs.command({
    command: 'read',
    describe: 'Reads all notes',
    builder: {
        title: {
            describe: 'Note Title',
            demandOption: true,
            type: 'string'
        }
    },
    handler(argv) {
        notes.readNote(argv.title);
    }
})

yargs.parse();


// const command = process.argv[2];

// if ( command === 'add' ) {
//     console.log(chalk.blue('Adding Note...'));
// } else if ( command === 'remove' ) {
//     console.log(chalk.red('Removing Note...'))
// }

