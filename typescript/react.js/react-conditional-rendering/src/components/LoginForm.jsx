import React from 'react'
import Input from './Input'
import Button from './Button'

const LoginForm = () => {
  return (
    <form className="form">
      <Input
        type="text"
        placeholder="Username"
      />
      <Input
        type="password"
        placeholder="Password"
      />
      <Button />
    </form>
  )
}

export default LoginForm
