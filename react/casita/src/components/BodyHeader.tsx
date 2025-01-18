import Header from './Header'
import Searchbar from './Searchbar'
import AccountCircleIcon from '@mui/icons-material/AccountCircle'
import NotificationImportantIcon from '@mui/icons-material/NotificationImportant'

function BodyHeader() {
  return (
    <div className='body-header'>
      <Header />
      <Searchbar />
      <NotificationImportantIcon
        fontSize='large'
        // className='notification-icon'
      />
      <AccountCircleIcon fontSize='large' />
    </div>
  )
}

export default BodyHeader
