import React from 'react'

function Card() {
  return (
    <div className="card">
      <img
        className="card--photo"
        src="./images/wedding-photography.png"
        alt="wedding-photo"
      />
      <p>
        <span className='card--star'>★</span> 5.0 <span className='card--origin'>(30) •USA </span>
      </p>
      <p>Learn wedding photography</p>
      <p>
        <span className='card--price'>From $125</span> / person
      </p>
    </div>
  )
}

export default Card
