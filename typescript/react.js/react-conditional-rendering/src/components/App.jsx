import React from 'react'
import RegisterForm from './RegisterForm'
import LoginForm from './LoginForm'

const userIsRegistered = true

const App = () => {
  return (
    <div className="container">
      {userIsRegistered ? <LoginForm /> : <RegisterForm />}
    </div>
  )
}

export default App
