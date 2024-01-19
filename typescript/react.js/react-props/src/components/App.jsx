import React from 'react'
import Card from './Card'
import Header from './Header'
import contacts from '../contacts'
import Avatar from './Avatar'

const createCard = contact => {
  return (
    <Card
      key={contact.id}
      id={contact.id}
      name={contact.name}
      img={contact.img}
      tel={contact.tel}
      email={contact.email}
    />
  )
}

const App = () => {
  return (
    <div>
      <Header />
      <Avatar img="https://upload.wikimedia.org/wikipedia/commons/thumb/7/77/20180602_N%C3%BCrnberg_Rock_im_Park_A_Perfect_Circle_0029_%28cropped%29.jpg/1024px-20180602_N%C3%BCrnberg_Rock_im_Park_A_Perfect_Circle_0029_%28cropped%29.jpg" />
      {contacts.map(createCard)}
    </div>
  )
}

export default App
