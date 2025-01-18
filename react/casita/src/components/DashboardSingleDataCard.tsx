import AccountBalanceIcon from '@mui/icons-material/AccountBalance'

const DashboardSingleCard = () => {
  return (
    <div className='dashboard-single-card'>
      <div className='dashboard-single-card-title'>
        <AccountBalanceIcon />
        <div>My Balance</div>
      </div>
      <div className='dashboard-single-card-content'>
        <div className='card-amount'>$5.000.000,00</div>
        <div className='card-label'>+33%</div>
      </div>
    </div>
  )
}

export default DashboardSingleCard
