import React, { useState } from 'react'
import Header from './Header'
import Footer from './Footer'
import CreateArea from './CreateArea'
import Note from './Note'

const App = () => {
  const [notes, setNotes] = useState([])

  const addNew = note => {
    setNotes(prevNotes => {
      return [...prevNotes, note]
    })
  }

  const deleteNote = id => {
    setNotes(prevNotes => {
      return prevNotes.filter((_elem, i) => {
        return i !== id
      })
    })
  }

  return (
    <div>
      <Header />
      <CreateArea onNew={addNew} />
      {notes.map((note, i) => (
        <Note
          key={i}
          title={note.title}
          content={note.content}
          onDelete={deleteNote}
          id={i}
        />
      ))}
      <Footer />
    </div>
  )
}

export default App
