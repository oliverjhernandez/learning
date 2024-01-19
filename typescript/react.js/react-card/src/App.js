import React from 'react'
import About from './components/About'
import Interests from './components/Interests'
import Info from './components/Info'
import Photo from './components/Photo'

function App() {
  return (
    <div>
      <div className="card">
        <Photo />
        <div className="card-info">
          <Info />
          <About />
          <Interests />
        </div>
      </div>
    </div>
  )
}

export default App
