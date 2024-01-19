import React from 'react'
import emojipedia from '../emojipedia'
import Entry from './Entry'
import Title from './Title'

const createEntry = (emoji) => {
  return (
    <Entry
      key={emoji.id}
      emoji={emoji.emoji}
      name={emoji.name}
      meaning={emoji.meaning}
    />
  )
}

const App = () => {
  return (
    <div>
      <Title />
      <dl className="dictionary">
        {emojipedia.map(createEntry)}
      </dl>
    </div>
  )
}

export default App
