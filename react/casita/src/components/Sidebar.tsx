import Button from './Button'
import '../styles/sidebar.css'
import Logo from './Logo'
import {
  CurrencyBitcoin,
  CreditScore,
  Logout,
  AccountBalance,
  AttachMoney,
  Construction,
  StackedLineChart,
} from '@mui/icons-material'

const testOnClick = () => void {}

function Sidebar() {
  return (
    <div className='sidebar'>
      <Logo />
      <Button name='Dashboard' icon={StackedLineChart} onClick={testOnClick} />
      <Button name='Expenses' icon={AttachMoney} onClick={testOnClick} />
      <Button name='Investments' icon={CurrencyBitcoin} onClick={testOnClick} />
      <Button name='Credits' icon={CreditScore} onClick={testOnClick} />
      <Button name='Budgets' icon={AccountBalance} onClick={testOnClick} />
      <Button name='Settings' icon={Construction} onClick={testOnClick} />
      <Button name='Logout' icon={Logout} onClick={testOnClick} />
    </div>
  )
}

export default Sidebar
