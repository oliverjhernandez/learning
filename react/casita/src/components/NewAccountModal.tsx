import { CloseTwoTone } from '@mui/icons-material'
import Button from './Button'

type NewAccountModalProps = {
  isOpen: boolean
  onClose: () => void
}

const NewAccountModal: React.FC<NewAccountModalProps> = ({
  isOpen,
  onClose,
}) => {
  if (!isOpen) {
    return null
  }
  console.log('Help!!')

  return (
    <div className='bg-seconday border border-alpha-black-25 rounded-2xl max-w-[580px] w-min-content shadow-xs h-fit '>
      <div>New Account</div>
      <Button name='Close' icon={CloseTwoTone} onClick={onClose} />
    </div>
  )
}

export default NewAccountModal
