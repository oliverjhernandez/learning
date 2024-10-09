import '../styles/button.css'
import { SvgIconTypeMap } from '@mui/material'
import { OverridableComponent } from '@mui/material/OverridableComponent'

type ButtonProps = {
  name: string
  icon: OverridableComponent<SvgIconTypeMap<{}, 'svg'>>
  onClick: () => void
}

const Button: React.FC<ButtonProps> = ({
  name,
  icon: Icon,
  onClick,
}: ButtonProps) => {
  return (
    <button className='btn' onClick={onClick}>
      <div className='flex flex-row'>
        <div className='btn-item'>
          <Icon />
        </div>
        <div className='btn-item'>{name}</div>
      </div>
    </button>
  )
}

export default Button
