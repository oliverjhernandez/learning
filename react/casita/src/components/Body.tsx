import { PlusOneRounded } from '@mui/icons-material'
import Button from './Button'
import NewAccountModal from './NewAccountModal'
import { useState } from 'react'

function Body() {
  const [isModalOpen, setIsModalOpen] = useState(false)

  const openNewAccountModal = () => {
    setIsModalOpen(true)
  }

  const closeNewAccountModal = () => {
    setIsModalOpen(false)
  }

  return (
    <div>
      <Button
        name='New Account'
        icon={PlusOneRounded}
        onClick={openNewAccountModal}
      />
      <NewAccountModal isOpen={isModalOpen} onClose={closeNewAccountModal} />
    </div>
  )
}

export default Body
