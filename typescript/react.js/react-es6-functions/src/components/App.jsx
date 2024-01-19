import React from 'react'
import emojipedia from '../emojipedia'

const App = () => {
  return <p>Hello</p>
}

export default App

const msgArray = emojipedia.map((term) => {
  return term.meaning.slice(0, 100)
})

console.log(msgArray)
