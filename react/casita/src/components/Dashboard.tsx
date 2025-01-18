import BodyDashboardBankList from './DashboardBankList'
import DashboardContainer from './DashboardContainer'
import DashboardExpenseGraph from './DashboardExpenseGraph'
import DashboardSavings from './DashboardSavings'
import DashboardSingleDataCard from './DashboardSingleDataCard'

function BodyDashboard() {
  return (
    <div className='dashboard'>
      <DashboardContainer className='dashboard-container-2x2' color='white'>
        <DashboardSingleDataCard />
        <DashboardSingleDataCard />
        <DashboardSingleDataCard />
        <DashboardSingleDataCard />
      </DashboardContainer>
      <DashboardContainer className='dashboard-container-2x1' color='pink'>
        <BodyDashboardBankList />
        <BodyDashboardBankList />
      </DashboardContainer>
      <DashboardContainer
        className='dashboard-container-1x1'
        color='lightgreen'
      >
        <DashboardExpenseGraph />
      </DashboardContainer>
      <DashboardContainer
        className='dashboard-container-1x2'
        color='lightyellow'
      >
        <DashboardSavings />
      </DashboardContainer>
    </div>
  )
}

export default BodyDashboard
