import Button from './Button'
import '../styles/sidebar.css'
import Logo from './Logo'
import { AccountBox, Dashboard, Paid } from '@mui/icons-material'

const testOnClick = () => void {}

function Sidebar() {
  return (
    <div className='sidebar'>
      <Logo />
      <Button name='Dashboard' icon={Dashboard} onClick={testOnClick} />
      <Button name='Accounts' icon={AccountBox} onClick={testOnClick} />
      <Button name='Transactions' icon={Paid} onClick={testOnClick} />
    </div>
  )
}

export default Sidebar
