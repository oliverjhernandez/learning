import Button from './Button'
import '../styles/sidebar.css'
import Logo from './Logo'
import { AccountBox, Dashboard, Paid } from '@mui/icons-material'

function Sidebar() {
  return (
    <div className='sidebar'>
      <Logo />
      <Button name='Dashboard' icon={Dashboard} />
      <Button name='Accounts' icon={AccountBox} />
      <Button name='Transactions' icon={Paid} />
    </div>
  )
}

export default Sidebar
